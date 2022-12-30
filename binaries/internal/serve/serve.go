package serve

import (
	"net"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var Cmd = &cobra.Command{
	Use:   "serve",
	Short: "serve",
	Long:  `serve for everything`,
	Run:   doServe,
}

type Serve interface {
	serve()
}

func doServe(cmd *cobra.Command, args []string) {
	addr := "127.0.0.1:63109"
	if len(args) > 0 {
		addr = args[0]
	}
	StartServe(addr)
}

func StartServe(addr string) {
	s := grpc.NewServer()

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatalf("net.Listen err: %v", err)
	}

	err = s.Serve(lis)
	if err != nil {
		logrus.Fatalf("server.Serve err: %v", err)
	}
}
