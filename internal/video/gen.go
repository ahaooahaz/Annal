package video

import (
	"fmt"
	"image"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gocv.io/x/gocv"
)

var GenCmd = &cobra.Command{
	Use:   "gen",
	Short: "gen video from single image",
	Long:  `gen video from single image, only support image(.jpeg) to video(.avi)`,
	Run: func(cmd *cobra.Command, args []string) {
		images, err := cmd.Flags().GetStringSlice("images")
		if err != nil {
			fmt.Printf("get images error:[%v]", err)
			return
		}

		output, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Printf("get output err:[%v]", err)
			return
		}

		fps, err := cmd.Flags().GetInt("fps")
		if err != nil {
			fmt.Printf("get fps error:[%v]", err.Error())
			return
		}

		N, err := cmd.Flags().GetUint64("N")
		if err != nil {
			fmt.Printf("get N error:[%v]", err)
			return
		}

		width, err := cmd.Flags().GetInt("width")
		if err != nil {
			fmt.Printf("get width error:[%v]", err)
			return
		}

		height, err := cmd.Flags().GetInt("height")
		if err != nil {
			fmt.Printf("get height error:[%v]", err)
			return
		}

		err = GenerateVideoFromImage(images, output, fps, width, height, N)
		if err != nil {
			fmt.Printf("error:[%v]\n", err.Error())
		}
	},
}

func init() {
	GenCmd.Flags().StringSliceP("image", "i", []string{}, "source image")
	GenCmd.Flags().StringP("output", "o", "annal.avi", "output video path")
	GenCmd.Flags().IntP("fps", "f", 25, "video fps")
	GenCmd.Flags().Uint64P("N", "", 25*60, "count of video frames")
	GenCmd.Flags().IntP("width", "w", 1920, "video width")
	GenCmd.Flags().IntP("height", "h", 1080, "video height")
}

func GenerateVideoFromImage(images []string, output string, fps int, width, height int, N uint64) (err error) {
	fourcc := "MP42"
	videoWriter, err := gocv.VideoWriterFile(output, fourcc, float64(fps), width, height, true)
	if err != nil {
		return
	}
	defer videoWriter.Close()

	if !videoWriter.IsOpened() {
		return fmt.Errorf("open video file failed")
	}

	mats := []*gocv.Mat{}
	for _, img := range images {
		mat := gocv.IMRead(img, gocv.IMReadColor)
		defer mat.Close()
		gocv.Resize(mat, &mat, image.Point{X: width, Y: height}, 0, 0, gocv.InterpolationArea)
		mats = append(mats, &mat)
	}

	for i := N * 0; i < N; i++ {
		err = videoWriter.Write(*mats[i%uint64(len(mats))])
		if err != nil {
			log.Errorf("write mat failed, error: %v", err.Error())
			return
		}
	}

	return
}
