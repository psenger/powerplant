package powerplant

type PlantType string

const (
	hydro PlantType = "hydro"
	wind PlantType = "wind"
	solar PlantType = "solar"
)

type PlantStatus string

const (
	active PlantStatus = "active"
	inactive PlantStatus = "inactive"
	unavailable PlantStatus = "unavailable"
)

type PowerPlant struct {
	plantType PlantType
	capacity float64
	status PlantStatus

}