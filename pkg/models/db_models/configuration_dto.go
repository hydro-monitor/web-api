package db_models

type ConfigurationDTO struct {
	NodeId        string
	Configuration string
}

func (c *ConfigurationDTO) GetColumns() []string {
	return nil
}
