package models

import (
	"crypto/md5"
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"time"

	"mime/multipart"

	"golang.org/x/image/bmp"

	"github.com/astaxie/beego/orm"
	"github.com/nfnt/resize"
)

type Image struct {
	ImageID       string   `orm:"size(40);pk"`
	UserID        string   `orm:"size(25)"`
	ImageName     string   `orm:"size(40)"`
	ThumbnailName string   `orm:"size(50)"`
	Path          string   `orm:"size(255)"`
	Year          int      `orm:"size(6)"`
	Month         int      `orm:"size(6)"`
	Day           int      `orm:"size(6)"`
	Suffix        string   `orm:"size(6)"`
	ImageFile     *os.File `orm:"-"`
	ThumbnailFile *os.File `orm:"-"`
}

func (img *Image) TableIndex() [][]string {
	return [][]string{
		[]string{"UserID", "ImageID"},
	}
}

type ImageDate struct {
	Year  int `orm:"size(6)"`
	Month int `orm:"size(6)"`
}

type imageDateSorter []ImageDate

func (ids imageDateSorter) Len() int {
	return len(ids)
}

func (ids imageDateSorter) Swap(i, j int) {
	ids[i], ids[j] = ids[j], ids[i]
}

func (ids imageDateSorter) Less(i, j int) bool {
	if ids[i].Year < ids[j].Year {
		return true
	} else if ids[i].Year == ids[j].Year {
		if ids[i].Month < ids[j].Month {
			return true
		}
	}
	return false
}

func duplicate(a []ImageDate) (ret []ImageDate) {
	for index := 0; index < len(a); index++ {
		if index > 0 && reflect.DeepEqual(a[index-1], a[index]) {
			continue
		}
		ret = append(ret, a[index])
	}
	return ret
}

var PWD string

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	PWD = strings.Replace(dir, "\\", "/", -1) + "/static"
}

func makeDirAll(img *Image) error {
	return os.MkdirAll(PWD+img.Path, os.ModeDir)
}

func scaleImage(in io.Reader, out io.Writer, width, height, quality int) error {
	origin, fm, err := image.Decode(in)
	if err != nil {
		fmt.Println("error: " + err.Error())
		return err
	}
	if width == 0 || height == 0 {
		width = origin.Bounds().Max.X
		height = origin.Bounds().Max.Y
	}
	if quality == 0 {
		quality = 100
	}
	canvas := resize.Thumbnail(uint(width), uint(height), origin, resize.Lanczos3)

	switch fm {
	case "jpeg":
		return jpeg.Encode(out, canvas, &jpeg.Options{quality})
	case "png":
		return png.Encode(out, canvas)
	case "gif":
		return gif.Encode(out, canvas, &gif.Options{})
	case "bmp":
		return bmp.Encode(out, canvas)
	default:
		return errors.New("ERROR FORMAT")
	}
	return nil
}

func CreateImage(img *Image, src multipart.File) error {
	defer src.Close()
	var err error
	if err = makeDirAll(img); err != nil {
		return err
	}

	if img.ImageFile, err = os.Create(PWD + img.Path + img.ImageName + img.Suffix); err != nil {
		return err
	}
	defer img.ImageFile.Close()

	if img.ThumbnailFile, err = os.Create(PWD + img.Path + img.ThumbnailName + img.Suffix); err != nil {
		return err
	}
	defer img.ThumbnailFile.Close()

	if _, err = io.Copy(img.ImageFile, src); err != nil {
		return err
	}

	if _, err := src.Seek(0, os.SEEK_SET); err != nil {
		return err
	}

	return scaleImage(src, img.ThumbnailFile, 200, 200, 100)
}

func RemoveImage(img *Image) error {
	if err := os.Remove(PWD + img.Path + img.ImageName + img.Suffix); err != nil {
		return err
	}

	return os.Remove(PWD + img.Path + img.ThumbnailName + img.Suffix)
}

func GenImageID(img *Image) {
	img.ImageID = fmt.Sprintf("%x", md5.Sum([]byte(img.Path+img.ImageName)))
	fmt.Println("imageID: " + img.ImageID)
}

func GenImageName(img *Image, header *multipart.FileHeader, now time.Time) {
	img.Suffix = strings.ToLower(filepath.Ext(header.Filename))
	img.ImageName = fmt.Sprintf("%x", md5.Sum([]byte(header.Filename+img.UserID+now.String())))
	img.ThumbnailName = img.ImageName + "_thumbnail"
	img.Path = fmt.Sprintf("/%d/%d/%d/", img.Year, img.Month, img.Day)
}

func GenTime(img *Image, now time.Time) {
	img.Year = now.Year()
	img.Month = int(now.Month())
	img.Day = now.Day()
}

func AddOneImage(header *multipart.FileHeader, UserID string) (*Image, error) {
	img := &Image{UserID: UserID}
	now := time.Now()
	GenTime(img, now)
	GenImageName(img, header, now)
	GenImageID(img)

	f, err := header.Open()
	if err != nil {
		return nil, err
	}

	if err := CreateImage(img, f); err != nil {
		RemoveImage(img)
		return nil, err
	}

	if _, err := O.Insert(img); err != nil {
		RemoveImage(img)
		return nil, err
	}
	return img, nil
}

func GetAllImage(userID string) ([]*Image, error) {
	var res []*Image
	qs := O.QueryTable("image")
	_, err := qs.Filter("user_i_d", userID).All(&res)
	return res, err
}

func GetSomeImages(userID string, year, month int) ([]*Image, error) {
	var res []*Image
	qs := O.QueryTable("image")
	_, err := qs.Filter("user_i_d", userID).Filter("year", year).Filter("month", month).All(&res)
	return res, err
}

func GetImageDate(userID string) ([]ImageDate, error) {
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("image.year", "image.month").From("image").Where("image.user_i_d = " + userID)
	sqlStr := qb.String()
	var imgDates []ImageDate

	O.Raw(sqlStr).QueryRows(&imgDates)

	sort.Sort(imageDateSorter(imgDates))

	res := duplicate(imgDates)

	return res, nil
}

func DeleteImage(ImageID, userID string) error {
	img := &Image{ImageID: ImageID, UserID: userID}
	if err := O.Read(img); err != nil {
		return err
	}

	if _, err := O.Delete(img); err != nil {
		return err
	}

	return RemoveImage(img)
}
