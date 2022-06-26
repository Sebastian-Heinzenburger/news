package main

import (
	"time"
)

type Regional struct {
	News []struct {
		SophoraId   string `json:"sophoraId"`
		ExternalId  string `json:"externalId"`
		Title       string `json:"title"`
		TeaserImage struct {
			Title             string `json:"title"`
			Alttext           string `json:"alttext"`
			PreferredVariants string `json:"preferredVariants"`
			Type              string `json:"type"`
			Videowebl         struct {
				Imageurl string `json:"imageurl"`
			} `json:"videowebl"`
			Portraetgrossplus8x9 struct {
				Imageurl string `json:"imageurl"`
			} `json:"portraetgrossplus8x9"`
			Videowebm struct {
				Imageurl string `json:"imageurl"`
			} `json:"videowebm"`
			Videowebs struct {
				Imageurl string `json:"imageurl"`
			} `json:"videowebs"`
			Portraetgross8x9 struct {
				Imageurl string `json:"imageurl"`
			} `json:"portraetgross8x9"`
		} `json:"teaserImage"`
		Date     time.Time `json:"date"`
		Tracking []struct {
			Sid  string `json:"sid"`
			Src  string `json:"src"`
			Ctp  string `json:"ctp"`
			Pdt  string `json:"pdt"`
			Otp  string `json:"otp"`
			Cid  string `json:"cid"`
			Pti  string `json:"pti"`
			Type string `json:"type"`
		} `json:"tracking"`
		Tags []struct {
			Tag string `json:"tag"`
		} `json:"tags"`
		UpdateCheckUrl string `json:"updateCheckUrl"`
		RegionId       int    `json:"regionId"`
		RegionIds      []int  `json:"regionIds"`
		Details        string `json:"details"`
		Detailsweb     string `json:"detailsweb"`
		ShareURL       string `json:"shareURL"`
		Topline        string `json:"topline"`
		FirstSentence  string `json:"firstSentence"`
		Geotags        []struct {
			Tag string `json:"tag"`
		} `json:"geotags"`
		BrandingImage struct {
			Title     string `json:"title"`
			Copyright string `json:"copyright"`
			Alttext   string `json:"alttext"`
			Type      string `json:"type"`
			Videowebm struct {
				Imageurl string `json:"imageurl"`
			} `json:"videowebm"`
			Videowebs struct {
				Imageurl string `json:"imageurl"`
			} `json:"videowebs"`
		} `json:"brandingImage"`
		Type string `json:"type"`
	} `json:"news"`
	Regional            []interface{} `json:"regional"`
	NewStoriesCountLink string        `json:"newStoriesCountLink"`
	NextPage            string        `json:"nextPage"`
	Type                string        `json:"type"`
}

type News struct {
	News []struct {
		SophoraId   string      `json:"sophoraId"`
		ExternalId  interface{} `json:"externalId"`
		Title       string      `json:"title"`
		TeaserImage struct {
			Title             string `json:"title"`
			Copyright         string `json:"copyright"`
			Alttext           string `json:"alttext"`
			PreferredVariants string `json:"preferredVariants"`
			Type              string `json:"type"`
			Videowebl         struct {
				Imageurl string `json:"imageurl"`
			} `json:"videowebl"`
			Portraetgrossplus8x9 struct {
				Imageurl string `json:"imageurl"`
			} `json:"portraetgrossplus8x9"`
			Videowebm struct {
				Imageurl string `json:"imageurl"`
			} `json:"videowebm"`
			Videowebs struct {
				Imageurl string `json:"imageurl"`
			} `json:"videowebs"`
			Portraetgross8x9 struct {
				Imageurl string `json:"imageurl"`
			} `json:"portraetgross8x9"`
		} `json:"teaserImage"`
		Content []struct {
			Value string `json:"value"`
			Type  string `json:"type"`
			Video struct {
				SophoraId   string `json:"sophoraId"`
				ExternalId  string `json:"externalId"`
				Title       string `json:"title"`
				TeaserImage struct {
					Title             string `json:"title"`
					Alttext           string `json:"alttext"`
					PreferredVariants string `json:"preferredVariants"`
					Type              string `json:"type"`
					Videowebl         struct {
						Imageurl string `json:"imageurl"`
					} `json:"videowebl"`
					Portraetgrossplus8x9 struct {
						Imageurl string `json:"imageurl"`
					} `json:"portraetgrossplus8x9"`
					Videowebm struct {
						Imageurl string `json:"imageurl"`
					} `json:"videowebm"`
					Videowebs struct {
						Imageurl string `json:"imageurl"`
					} `json:"videowebs"`
					Portraetgross8x9 struct {
						Imageurl string `json:"imageurl"`
					} `json:"portraetgross8x9"`
				} `json:"teaserImage"`
				Date     time.Time `json:"date"`
				Tracking []struct {
					Sid         string `json:"sid"`
					Src         string `json:"src"`
					Ctp         string `json:"ctp"`
					Pdt         string `json:"pdt"`
					Otp         string `json:"otp"`
					Cid         string `json:"cid"`
					Pti         string `json:"pti"`
					Type        string `json:"type"`
					Assetid     string `json:"assetid"`
					Program     string `json:"program"`
					Title       string `json:"title"`
					Length      string `json:"length"`
					C2          string `json:"c2"`
					C5          string `json:"c5"`
					C7          string `json:"c7"`
					C8          string `json:"c8"`
					C9          string `json:"c9"`
					C10         string `json:"c10"`
					C12         string `json:"c12"`
					C16         string `json:"c16"`
					C18         string `json:"c18"`
					TypeNielsen string `json:"type_nielsen"`
				} `json:"tracking"`
				Tags           []interface{} `json:"tags"`
				UpdateCheckUrl string        `json:"updateCheckUrl"`
				Streams        struct {
					H264m             string `json:"h264m"`
					H264s             string `json:"h264s"`
					H264xl            string `json:"h264xl"`
					Adaptivestreaming string `json:"adaptivestreaming"`
				} `json:"streams"`
				Copyright string `json:"copyright"`
				Type      string `json:"type"`
				Alttext   string `json:"alttext"`
			} `json:"video"`
			Box struct {
				Title    string `json:"title"`
				Subtitle string `json:"subtitle"`
				Images   struct {
					Title             string `json:"title"`
					Copyright         string `json:"copyright"`
					Alttext           string `json:"alttext"`
					PreferredVariants string `json:"preferredVariants"`
					Type              string `json:"type"`
					Videowebl         struct {
						Imageurl string `json:"imageurl"`
					} `json:"videowebl"`
					Portraetgrossplus8x9 struct {
						Imageurl string `json:"imageurl"`
					} `json:"portraetgrossplus8x9"`
					Videowebm struct {
						Imageurl string `json:"imageurl"`
					} `json:"videowebm"`
					Videowebs struct {
						Imageurl string `json:"imageurl"`
					} `json:"videowebs"`
					Portraetgross8x9 struct {
						Imageurl string `json:"imageurl"`
					} `json:"portraetgross8x9"`
				} `json:"images"`
				Text   string `json:"text"`
				Link   string `json:"link"`
				Source string `json:"source"`
			} `json:"box"`
			Gallery []struct {
				Title             string `json:"title"`
				Alttext           string `json:"alttext"`
				PreferredVariants string `json:"preferredVariants"`
				Type              string `json:"type"`
				Videowebl         struct {
					Imageurl string `json:"imageurl"`
				} `json:"videowebl"`
				Portraetgrossplus8x9 struct {
					Imageurl string `json:"imageurl"`
				} `json:"portraetgrossplus8x9"`
				Videowebm struct {
					Imageurl string `json:"imageurl"`
				} `json:"videowebm"`
				Videowebs struct {
					Imageurl string `json:"imageurl"`
				} `json:"videowebs"`
				Portraetgross8x9 struct {
					Imageurl string `json:"imageurl"`
				} `json:"portraetgross8x9"`
			} `json:"gallery"`
			Tracking []struct {
				Sid  string `json:"sid"`
				Src  string `json:"src"`
				Ctp  string `json:"ctp"`
				Pdt  string `json:"pdt"`
				Otp  string `json:"otp"`
				Cid  string `json:"cid"`
				Pti  string `json:"pti"`
				Type string `json:"type"`
			} `json:"tracking"`
			Related []struct {
				TeaserImage struct {
					Title             string `json:"title"`
					Copyright         string `json:"copyright"`
					Alttext           string `json:"alttext"`
					PreferredVariants string `json:"preferredVariants"`
					Type              string `json:"type"`
					Videowebl         struct {
						Imageurl string `json:"imageurl"`
					} `json:"videowebl"`
					Portraetgrossplus8x9 struct {
						Imageurl string `json:"imageurl"`
					} `json:"portraetgrossplus8x9"`
					Videowebm struct {
						Imageurl string `json:"imageurl"`
					} `json:"videowebm"`
					Videowebs struct {
						Imageurl string `json:"imageurl"`
					} `json:"videowebs"`
					Portraetgross8x9 struct {
						Imageurl string `json:"imageurl"`
					} `json:"portraetgross8x9"`
				} `json:"teaserImage"`
				Date       time.Time `json:"date"`
				SophoraId  string    `json:"sophoraId"`
				ExternalId string    `json:"externalId"`
				Topline    string    `json:"topline"`
				Title      string    `json:"title"`
				Details    string    `json:"details"`
				Type       string    `json:"type"`
				Detailsweb string    `json:"detailsweb"`
				Streams    struct {
					H264m             string `json:"h264m"`
					H264s             string `json:"h264s"`
					H264xl            string `json:"h264xl"`
					Adaptivestreaming string `json:"adaptivestreaming"`
				} `json:"streams"`
			} `json:"related"`
		} `json:"content"`
		Date     time.Time `json:"date"`
		Tracking []struct {
			Sid  string `json:"sid"`
			Src  string `json:"src"`
			Ctp  string `json:"ctp"`
			Pdt  string `json:"pdt"`
			Otp  string `json:"otp"`
			Cid  string `json:"cid"`
			Pti  string `json:"pti"`
			Type string `json:"type"`
		} `json:"tracking"`
		Tags []struct {
			Tag string `json:"tag"`
		} `json:"tags"`
		UpdateCheckUrl string `json:"updateCheckUrl"`
		RegionId       int    `json:"regionId"`
		Images         []struct {
			Title             string `json:"title"`
			Alttext           string `json:"alttext"`
			PreferredVariants string `json:"preferredVariants"`
			Type              string `json:"type"`
			Videowebl         struct {
				Imageurl string `json:"imageurl"`
			} `json:"videowebl"`
			Portraetgrossplus8x9 struct {
				Imageurl string `json:"imageurl"`
			} `json:"portraetgrossplus8x9"`
			Videowebm struct {
				Imageurl string `json:"imageurl"`
			} `json:"videowebm"`
			Videowebs struct {
				Imageurl string `json:"imageurl"`
			} `json:"videowebs"`
			Portraetgross8x9 struct {
				Imageurl string `json:"imageurl"`
			} `json:"portraetgross8x9"`
		} `json:"images"`
		Details       string `json:"details"`
		Detailsweb    string `json:"detailsweb"`
		ShareURL      string `json:"shareURL"`
		Topline       string `json:"topline"`
		FirstSentence string `json:"firstSentence"`
		Geotags       []struct {
			Tag string `json:"tag"`
		} `json:"geotags"`
		Crop struct {
			Id                     string        `json:"id"`
			Type                   string        `json:"type"`
			CroppingApiVersion     string        `json:"croppingApiVersion"`
			CroppingUIVersion      string        `json:"croppingUIVersion"`
			CroppingServiceVersion string        `json:"croppingServiceVersion"`
			NoSound                bool          `json:"noSound"`
			VideoSrc               string        `json:"videoSrc"`
			ImageSrc               []interface{} `json:"imageSrc"`
			ImageNames             []interface{} `json:"imageNames"`
			HeaderText             string        `json:"headerText"`
			MainTexts              []string      `json:"mainTexts"`
			SceneSrcTexts          []string      `json:"sceneSrcTexts"`
			CameraMoves            []struct {
				Point1X   float64 `json:"point1X"`
				Point1Y   float64 `json:"point1Y"`
				Point2X   float64 `json:"point2X"`
				Point2Y   float64 `json:"point2Y"`
				StartLeft int     `json:"startLeft"`
				EndLeft   int     `json:"endLeft"`
				Duration  int     `json:"duration"`
			} `json:"cameraMoves"`
			Events []interface{} `json:"events"`
		} `json:"crop"`
		Ressort  string `json:"ressort"`
		Comments string `json:"comments"`
		Type     string `json:"type"`
		Streams  struct {
			H264m             string `json:"h264m"`
			H264s             string `json:"h264s"`
			H264xl            string `json:"h264xl"`
			Adaptivestreaming string `json:"adaptivestreaming"`
		} `json:"streams"`
		BreakingNews bool `json:"breakingNews"`
		FirstFrame   struct {
			Title             string `json:"title"`
			Alttext           string `json:"alttext"`
			PreferredVariants string `json:"preferredVariants"`
			Type              string `json:"type"`
			Videowebl         struct {
				Imageurl string `json:"imageurl"`
			} `json:"videowebl"`
			Portraetgrossplus8x9 struct {
				Imageurl string `json:"imageurl"`
			} `json:"portraetgrossplus8x9"`
			Videowebm struct {
				Imageurl string `json:"imageurl"`
			} `json:"videowebm"`
			Videowebs struct {
				Imageurl string `json:"imageurl"`
			} `json:"videowebs"`
			Portraetgross8x9 struct {
				Imageurl string `json:"imageurl"`
			} `json:"portraetgross8x9"`
		} `json:"firstFrame"`
		RegionIds     []int `json:"regionIds"`
		BrandingImage struct {
			Title     string `json:"title"`
			Copyright string `json:"copyright"`
			Alttext   string `json:"alttext"`
			Type      string `json:"type"`
			Videowebm struct {
				Imageurl string `json:"imageurl"`
			} `json:"videowebm"`
			Videowebs struct {
				Imageurl string `json:"imageurl"`
			} `json:"videowebs"`
		} `json:"brandingImage"`
		Shorttext string `json:"shorttext"`
	} `json:"news"`
	Regional []struct {
		SophoraId   string `json:"sophoraId"`
		ExternalId  string `json:"externalId"`
		Title       string `json:"title"`
		TeaserImage struct {
			Title             string `json:"title"`
			Alttext           string `json:"alttext"`
			PreferredVariants string `json:"preferredVariants"`
			Type              string `json:"type"`
			Videowebl         struct {
				Imageurl string `json:"imageurl"`
			} `json:"videowebl"`
			Portraetgrossplus8x9 struct {
				Imageurl string `json:"imageurl"`
			} `json:"portraetgrossplus8x9"`
			Videowebm struct {
				Imageurl string `json:"imageurl"`
			} `json:"videowebm"`
			Videowebs struct {
				Imageurl string `json:"imageurl"`
			} `json:"videowebs"`
			Portraetgross8x9 struct {
				Imageurl string `json:"imageurl"`
			} `json:"portraetgross8x9"`
		} `json:"teaserImage"`
		Content []struct {
			Value string `json:"value"`
			Type  string `json:"type"`
			Video struct {
				SophoraId   string `json:"sophoraId"`
				ExternalId  string `json:"externalId"`
				Title       string `json:"title"`
				TeaserImage struct {
					Title             string `json:"title"`
					Alttext           string `json:"alttext"`
					PreferredVariants string `json:"preferredVariants"`
					Type              string `json:"type"`
					Videowebl         struct {
						Imageurl string `json:"imageurl"`
					} `json:"videowebl"`
					Portraetgrossplus8x9 struct {
						Imageurl string `json:"imageurl"`
					} `json:"portraetgrossplus8x9"`
					Videowebm struct {
						Imageurl string `json:"imageurl"`
					} `json:"videowebm"`
					Videowebs struct {
						Imageurl string `json:"imageurl"`
					} `json:"videowebs"`
					Portraetgross8x9 struct {
						Imageurl string `json:"imageurl"`
					} `json:"portraetgross8x9"`
				} `json:"teaserImage"`
				Date     time.Time `json:"date"`
				Tracking []struct {
					Sid  string `json:"sid"`
					Src  string `json:"src"`
					Ctp  string `json:"ctp"`
					Pdt  string `json:"pdt"`
					Otp  string `json:"otp"`
					Cid  string `json:"cid"`
					Pti  string `json:"pti"`
					Type string `json:"type"`
				} `json:"tracking"`
				Tags           []interface{} `json:"tags"`
				UpdateCheckUrl string        `json:"updateCheckUrl"`
				Streams        struct {
					H264xl            string `json:"h264xl"`
					Adaptivestreaming string `json:"adaptivestreaming"`
				} `json:"streams"`
				Copyright string `json:"copyright"`
				Type      string `json:"type"`
				Alttext   string `json:"alttext"`
			} `json:"video"`
			Box struct {
				Title    string `json:"title"`
				Subtitle string `json:"subtitle"`
				Text     string `json:"text"`
				Source   string `json:"source"`
			} `json:"box"`
		} `json:"content"`
		Date     time.Time `json:"date"`
		Tracking []struct {
			Sid  string `json:"sid"`
			Src  string `json:"src"`
			Ctp  string `json:"ctp"`
			Pdt  string `json:"pdt"`
			Otp  string `json:"otp"`
			Cid  string `json:"cid"`
			Pti  string `json:"pti"`
			Type string `json:"type"`
		} `json:"tracking"`
		Tags []struct {
			Tag string `json:"tag"`
		} `json:"tags"`
		UpdateCheckUrl string `json:"updateCheckUrl"`
		RegionId       int    `json:"regionId"`
		RegionIds      []int  `json:"regionIds"`
		Images         []struct {
			Title             string `json:"title"`
			Alttext           string `json:"alttext"`
			PreferredVariants string `json:"preferredVariants"`
			Type              string `json:"type"`
			Videowebl         struct {
				Imageurl string `json:"imageurl"`
			} `json:"videowebl"`
			Portraetgrossplus8x9 struct {
				Imageurl string `json:"imageurl"`
			} `json:"portraetgrossplus8x9"`
			Videowebm struct {
				Imageurl string `json:"imageurl"`
			} `json:"videowebm"`
			Videowebs struct {
				Imageurl string `json:"imageurl"`
			} `json:"videowebs"`
			Portraetgross8x9 struct {
				Imageurl string `json:"imageurl"`
			} `json:"portraetgross8x9"`
		} `json:"images"`
		Details       string `json:"details"`
		Detailsweb    string `json:"detailsweb"`
		ShareURL      string `json:"shareURL"`
		Topline       string `json:"topline"`
		FirstSentence string `json:"firstSentence"`
		Geotags       []struct {
			Tag string `json:"tag"`
		} `json:"geotags"`
		BrandingImage struct {
			Title     string `json:"title"`
			Copyright string `json:"copyright"`
			Alttext   string `json:"alttext"`
			Type      string `json:"type"`
			Videowebm struct {
				Imageurl string `json:"imageurl"`
			} `json:"videowebm"`
			Videowebs struct {
				Imageurl string `json:"imageurl"`
			} `json:"videowebs"`
		} `json:"brandingImage"`
		Type         string `json:"type"`
		BreakingNews bool   `json:"breakingNews"`
	} `json:"regional"`
	NewStoriesCountLink string `json:"newStoriesCountLink"`
	Type                string `json:"type"`
}
