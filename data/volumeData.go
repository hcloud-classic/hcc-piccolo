package data

import "hcc/piccolo/model"

// Query

// VolumeData : Data structure of volume
type VolumeData struct {
	Data struct {
		Volume model.Volume `json:"volume"`
	} `json:"data"`
}

// ListVolumeData : Data structure of list_volume
type ListVolumeData struct {
	Data struct {
		ListVolume []model.Volume `json:"list_volume"`
	} `json:"data"`
}

// AllVolumeData : Data structure of all_volume
type AllVolumeData struct {
	Data struct {
		AllVolume []model.Volume `json:"all_volume"`
	} `json:"data"`
}

// NumVolumeData : Data structure of num_volume
type NumVolumeData struct {
	Data struct {
		NumVolume model.VolumeNum `json:"num_volume"`
	} `json:"data"`
}

// VolumeAttatchmentData : Data structure of volume_attachment
type VolumeAttatchmentData struct {
	Data struct {
		VolumeAttatchment model.VolumeAttachment `json:"volume_attachment"`
	} `json:"data"`
}

// ListVolumeAttatchmentData : Data structure of list_volume_attachment
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

// Mutation

// CreateVolumeData : Data structure of create_volume
type CreateVolumeData struct {
	Data struct {
		Volume model.Volume `json:"create_volume"`
	} `json:"data"`
}

// UpdateVolumeData : Data structure of update_volume
type UpdateVolumeData struct {
	Data struct {
		Volume model.Volume `json:"update_volume"`
	} `json:"data"`
}

// DeleteVolumeData : Data structure of delete_volume
type DeleteVolumeData struct {
	Data struct {
		Volume model.Volume `json:"delete_volume"`
	} `json:"data"`
}

// CreateVolumeAttatchmentData : Data structure of create_volume_attachment
type CreateVolumeAttatchmentData struct {
	Data struct {
		VolumeAttachment model.VolumeAttachment `json:"create_volume_attachment"`
	} `json:"data"`
}

// UpdateVolumeAttatchmentData : Data structure of update_volume_attachment
type UpdateVolumeAttatchmentData struct {
	Data struct {
		VolumeAttachment model.VolumeAttachment `json:"update_volume_attachment"`
	} `json:"data"`
}

// DeleteVolumeAttatchmentData : Data structure of delete_volume_attachment
type DeleteVolumeAttatchmentData struct {
	Data struct {
		VolumeAttachment model.VolumeAttachment `json:"delete_volume_attachment"`
	} `json:"data"`
}
