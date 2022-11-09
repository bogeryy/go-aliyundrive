/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package aliyundrive

import (
	"context"
	"net/http"
)

// GetFile 获取文件信息
func (r *FileService) GetVideoPreviewPlayInfo(ctx context.Context, request *GetVideoPreviewPlayInfoReq) (*GetVideoPreviewPlayInfoResp, error) {
	request.Category = "live_transcoding"
	request.GetSubtitleInfo = true
	req := &RawRequestReq{
		Scope:  "File",
		API:    "GetVideoPreviewPlayInfo",
		Method: http.MethodPost,
		URL:    "https://api.aliyundrive.com/v2/file/get_video_preview_play_info",
		Body:   request,
	}
	resp := new(GetVideoPreviewPlayInfoResp)

	if _, err := r.cli.RawRequest(ctx, req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type GetVideoPreviewPlayInfoReq struct {
	DriveID         string `json:"drive_id"`
	FileID          string `json:"file_id"`
	Category        string `json:"category"`
	TemplateID      string `json:"template_id"`
	GetSubtitleInfo bool   `json:"get_subtitle_info"`
}

type GetVideoPreviewPlayInfoResp struct {
	DomainID             string `json:"domain_id"`
	DriveID              string `json:"drive_id"`
	FileID               string `json:"file_id"`
	VideoPreviewPlayInfo struct {
		Category string `json:"category"`
		Meta     struct {
			Duration            float64 `json:"duration"`
			Width               int     `json:"width"`
			Height              int     `json:"height"`
			LiveTranscodingMeta struct {
				TsSegment    int `json:"ts_segment"`
				TsTotalCount int `json:"ts_total_count"`
				TsPreCount   int `json:"ts_pre_count"`
			} `json:"live_transcoding_meta"`
		} `json:"meta"`
		LiveTranscodingTaskList []struct {
			TemplateID             string `json:"template_id"`
			TemplateName           string `json:"template_name"`
			TemplateWidth          int    `json:"template_width"`
			TemplateHeight         int    `json:"template_height"`
			Status                 string `json:"status"`
			Stage                  string `json:"stage"`
			URL                    string `json:"url"`
			KeepOriginalResolution bool   `json:"keep_original_resolution,omitempty"`
		} `json:"live_transcoding_task_list"`
	} `json:"video_preview_play_info"`
}
