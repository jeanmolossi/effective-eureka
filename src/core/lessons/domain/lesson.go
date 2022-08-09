package domain

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type lesson struct {
	sectionID    string
	lessonID     string
	index        uint16
	title        string
	description  string
	thumb        string
	videoPreview string
	video        string
	published    bool
	createdAt    time.Time
	updatedAt    time.Time
}

func NewLesson(sectionID, lessonID, title, description, thumb, videoPreview, video string, index uint16, published bool, createdAt, updatedAt *time.Time) Lesson {
	less := &lesson{
		sectionID:    sectionID,
		lessonID:     lessonID,
		index:        index,
		title:        title,
		description:  description,
		thumb:        thumb,
		videoPreview: videoPreview,
		video:        video,
		published:    published,
	}

	if createdAt != nil {
		less.createdAt = *createdAt
	}

	if updatedAt != nil {
		less.updatedAt = *updatedAt
	}

	return less
}

func (l *lesson) GetSectionID() string {
	return l.sectionID
}

func (l *lesson) GetLessonID() string {
	return l.lessonID
}

func (l *lesson) GetTitle() string {
	return l.title
}

func (l *lesson) GetDescription() string {
	return l.description
}

func (l *lesson) GetThumbnail() string {
	return l.thumb
}

func (l *lesson) GetVideoPreview() string {
	return l.videoPreview
}

func (l *lesson) GetVideo() string {
	if l.video == "" {
		return ""
	}

	return fmt.Sprintf("%s?%s", l.video, l.GenerateQueryValidation().Encode())
}

func (l *lesson) GetIndex() uint16 {
	return l.index
}

func (l *lesson) IsPublished() bool {
	return l.published
}

func (l *lesson) GetTimestamps() (createdAt, updatedAt time.Time) {
	return l.createdAt, l.updatedAt
}

func (l *lesson) SetSectionID(sectionID string) {
	l.sectionID = sectionID
}

func (l *lesson) SetLessonID(lessonID string) {
	l.lessonID = lessonID
}

func (l *lesson) SetTitle(title string) {
	l.title = title
}

func (l *lesson) SetDescription(description string) {
	l.description = description
}

func (l *lesson) SetThumbnail(thumbnail string) {
	l.thumb = thumbnail
}

func (l *lesson) SetVideoPreview(previewUrl string) {
	l.videoPreview = previewUrl
}

func (l *lesson) SetVideo(video string) {
	l.video = video
}

func (l *lesson) SetIndex(index uint16) {
	l.index = index
}

func (l *lesson) GenerateQueryValidation() url.Values {
	timestamp := time.Now()
	expiry := timestamp.Add(time.Minute * 10).Unix()
	message := fmt.Sprintf("%s:%d:%d", l.lessonID, timestamp.Unix(), expiry)

	hash := hmac.New(sha1.New, []byte("secret"))
	hash.Write([]byte(message))
	signature := hex.EncodeToString(hash.Sum(nil))

	return url.Values{
		"expiry": {fmt.Sprintf("%d", expiry)},
		"hash":   {signature},
	}
}

func (l *lesson) ValidateVideoLink(query url.Values) bool {
	expiryStr := query.Get("expiry")
	hashStr := query.Get("hash")

	if expiryStr == "" || hashStr == "" {
		return false
	}

	expiryUnix, err := strconv.ParseInt(expiryStr, 10, 64)
	if err != nil {
		return false
	}

	expiry := time.Unix(expiryUnix, 0)
	timestamp := expiry.Add(time.Minute * -10)
	message := fmt.Sprintf("%s:%d:%d", l.lessonID, timestamp.Unix(), expiryUnix)
	hash := hmac.New(sha1.New, []byte("secret"))
	hash.Write([]byte(message))
	signature := hex.EncodeToString(hash.Sum(nil))

	return signature == hashStr
}

func (l *lesson) Publish() {
	l.published = true
}

func (l *lesson) Unpublish() {
	l.published = false
}
