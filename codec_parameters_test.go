package astiav

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCodecParameters(t *testing.T) {
	fc, err := globalHelper.inputFormatContext("video.mp4")
	require.NoError(t, err)
	ss := fc.Streams()
	require.Len(t, ss, 2)
	s1 := ss[0]
	s2 := ss[1]

	cp1 := s1.CodecParameters()
	require.Equal(t, int64(441324), cp1.BitRate())
	require.Equal(t, ChromaLocationLeft, cp1.ChromaLocation())
	require.Equal(t, CodecIDH264, cp1.CodecID())
	require.Equal(t, CodecTag(0x31637661), cp1.CodecTag())
	require.Equal(t, ColorPrimariesUnspecified, cp1.ColorPrimaries())
	require.Equal(t, ColorRangeUnspecified, cp1.ColorRange())
	require.Equal(t, ColorSpaceUnspecified, cp1.ColorSpace())
	require.Equal(t, ColorTransferCharacteristicUnspecified, cp1.ColorTransferCharacteristic())
	require.Equal(t, 180, cp1.Height())
	require.Equal(t, Level(13), cp1.Level())
	require.Equal(t, MediaTypeVideo, cp1.MediaType())
	require.Equal(t, PixelFormatYuv420P, cp1.PixelFormat())
	require.Equal(t, ProfileH264ConstrainedBaseline, cp1.Profile())
	require.Equal(t, NewRational(1, 1), cp1.SampleAspectRatio())
	require.Equal(t, 320, cp1.Width())

	cp2 := s2.CodecParameters()
	require.Equal(t, int64(161052), cp2.BitRate())
	require.Equal(t, 2, cp2.Channels())
	require.True(t, cp2.ChannelLayout().Equal(ChannelLayoutStereo))
	require.Equal(t, CodecIDAac, cp2.CodecID())
	require.Equal(t, CodecTag(0x6134706d), cp2.CodecTag())
	require.Equal(t, 1024, cp2.FrameSize())
	require.Equal(t, MediaTypeAudio, cp2.MediaType())
	require.Equal(t, SampleFormatFltp, cp2.SampleFormat())
	require.Equal(t, 48000, cp2.SampleRate())

	cp3 := AllocCodecParameters()
	require.NotNil(t, cp3)
	defer cp3.Free()
	err = cp2.Copy(cp3)
	require.NoError(t, err)
	require.Equal(t, 2, cp3.Channels())

	cc4 := AllocCodecContext(nil)
	require.NotNil(t, cc4)
	defer cc4.Free()
	err = cp2.ToCodecContext(cc4)
	require.NoError(t, err)
	require.Equal(t, 2, cc4.Channels())

	cp5 := AllocCodecParameters()
	require.NotNil(t, cp5)
	defer cp5.Free()
	err = cp5.FromCodecContext(cc4)
	require.NoError(t, err)
	require.Equal(t, 2, cp5.Channels())

	cp6 := AllocCodecParameters()
	require.NotNil(t, cp6)
	defer cp6.Free()
	cp6.SetChannelLayout(ChannelLayout21)
	require.True(t, cp6.ChannelLayout().Equal(ChannelLayout21))
	defer cp6.Free()
	cp6.SetChannels(3)
	require.Equal(t, 3, cp6.Channels())
	cp6.SetCodecID(CodecIDRawvideo)
	require.Equal(t, CodecIDRawvideo, cp6.CodecID())
	cp6.SetCodecTag(CodecTag(2))
	require.Equal(t, CodecTag(2), cp6.CodecTag())
	cp6.SetColorRange(ColorRangeJpeg)
	require.Equal(t, ColorRangeJpeg, cp6.ColorRange())
	cp6.SetCodecType(MediaTypeAudio)
	require.Equal(t, MediaTypeAudio, cp6.CodecType())
	cp6.SetFrameSize(1)
	require.Equal(t, 1, cp6.FrameSize())
	cp6.SetHeight(1)
	require.Equal(t, 1, cp6.Height())
	cp1.SetLevel(16)
	require.Equal(t, Level(16), cp1.Level())
	cp6.SetMediaType(MediaTypeAudio)
	require.Equal(t, MediaTypeAudio, cp6.MediaType())
	cp1.SetProfile(ProfileH264Extended)
	require.Equal(t, ProfileH264Extended, cp1.Profile())
	cp6.SetPixelFormat(PixelFormat0Bgr)
	require.Equal(t, PixelFormat0Bgr, cp6.PixelFormat())
	cp6.SetSampleAspectRatio(NewRational(1, 2))
	require.Equal(t, NewRational(1, 2), cp6.SampleAspectRatio())
	cp6.SetSampleFormat(SampleFormatDbl)
	require.Equal(t, SampleFormatDbl, cp6.SampleFormat())
	cp6.SetSampleRate(4)
	require.Equal(t, 4, cp6.SampleRate())
	cp6.SetWidth(2)
	require.Equal(t, 2, cp6.Width())

	extraBytes := []byte{0, 0, 0, 1}
	require.NoError(t, cp6.SetExtraData(extraBytes))
	require.Equal(t, extraBytes, cp6.ExtraData())
}
