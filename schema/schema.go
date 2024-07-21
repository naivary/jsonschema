package schema


type property interface {
    // Set is checking if the marker is a valid option for the property
    // iff then the value will be set.
    Set(marker string, v any) error
}
