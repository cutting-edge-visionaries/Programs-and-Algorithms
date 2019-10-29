// This test file sends 100000 messages to committer
// and checks if the commit message was accepted.

package test

import (
	"context"
	mb "daad/protos/master"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"gotest.tools/assert"
	"testing"
	"time"
)

var (
	validationHost = flag.String("validator_host", "127.0.0.1", "ip address of the validator")
	committerHost  = flag.String("committer_host", "127.0.0.1", "ip address of the committer")
	validationPort = flag.Int("validator_port", 29000, "port on which validator will accept rpc calls")
	committerPort  = flag.Int("committer_port", 28000, "port on which committer will accept rpc calls")
)

func TestCommit(t *testing.T) {
	flag.Parse()

	// Create a GRPC client for validator
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	validationConn, err := grpc.Dial(fmt.Sprintf("%s:%d", *validationHost, *validationPort), opts...)

	// error must be nil
	assert.Equal(t, err, nil)
	validationClient := mb.NewValidationClient(validationConn)

	// Create a GRPC client for committer
	commitConn, err := grpc.Dial(fmt.Sprintf("%s:%d", *committerHost, *committerPort), opts...)

	// error must be nil
	assert.Equal(t, err, nil)
	commitClient := mb.NewCommitClient(commitConn)

	defer validationConn.Close()
	defer commitConn.Close()

	// Send non-validated messages to committer
	for i := int64(0); i < 10; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		response, err := commitClient.SubmitRequest(ctx, &mb.ValidationResponse{
			Msg:   "Message without validation",
			MsgId: i,
		})



		assert.Equal(t, err, nil)

		assert.Equal(t, response.ReturnValue, mb.CommitResponse_FAILURE)
	}

	// Send validated messages to committer
	for i := int64(10); i < 20; i++ {
		vctx, vcancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer vcancel()

		validatedMsg, err := validationClient.SubmitRequest(vctx, &mb.ValidationRequest{
			Msg:   "Sample message",
			MsgId: i,
		})



		assert.Equal(t, err, nil)

		cctx, ccancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer ccancel()

		response, err := commitClient.SubmitRequest(cctx, validatedMsg)



		assert.Equal(t, err, nil)

		assert.Equal(t, response.ReturnValue, mb.CommitResponse_SUCCESS)
	}
}
