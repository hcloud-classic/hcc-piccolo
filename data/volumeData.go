package data

import "hcc/piccolo/model"

type VolumeData struct {
	Data struct {
		Volume model.Volume `json:"volume"`
	} `json:"data"`
}

type ListVolumeData struct {
	Data struct {
		ListVolume []model.Volume `json:"list_volume"`
	} `json:"data"`
}

type AllVolumeData struct {
	Data struct {
		AllVolume []model.Volume `json:"all_volume"`
	} `json:"data"`
}

type NumVolumeData struct {
	Data struct {
		NumVolume model.VolumeNum `json:"num_volume"`
	} `json:"data"`
}

type VolumeAttatchmentData struct {
	Data struct {
		VolumeAttatchment model.VolumeAttachment `json:"volume_attachment"`
	} `json:"data"`
}

type ListVolumeAttatchmentData struct {
	Data struct {
		ListVolumeAttatchment []model.VolumeAttachment `json:"list_volume_attachment"`
	} `json:"data"`
}

// AllVolumeAttatchmentData : Data structure of all_volume_attachment
type AllVolumeAttatchmentData struct {
	Data struct {
		AllVolumeAttatchment []model.VolumeAttachment `json:"all_volume_attachment"`
	} `json:"data"`
}

type CreateVolumeData struct {
	Data struct {
		Volume model.Volume `json:"create_volume"`
	} `json:"data"`
}

type UpdateVolumeData struct {
	Data struct {
		Volume model.Volume `json:"update_volume"`
	} `json:"data"`
}

type DeleteVolumeData struct {
	Data struct {
		Volume model.Volume `json:"delete_volume"`
	} `json:"data"`
}

type CreateVolumeAttatchmentData struct {
	Data struct {
		VolumeAttachment model.VolumeAttachment `json:"create_volume_attachment"`
	} `json:"data"`
}

type UpdateVolumeAttatchmentData struct {
	Data struct {
		VolumeAttachment model.VolumeAttachment `json:"update_volume_attachment"`
	} `json:"data"`
}

type DeleteVolumeAttatchmentData struct {
	Data struct {
		VolumeAttachment model.VolumeAttachment `json:"delete_volume_attachment"`
	} `json:"data"`
}
