/*
 * Namf_Location
 *
 * AMF Location Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type SupportedGadShapes string

// List of SupportedGADShapes
const (
	SupportedGadShapes_POINT                      SupportedGadShapes = "POINT"
	SupportedGadShapes_POINT_UNCERTAINTY_CIRCLE   SupportedGadShapes = "POINT_UNCERTAINTY_CIRCLE"
	SupportedGadShapes_POINT_UNCERTAINTY_ELLIPSE  SupportedGadShapes = "POINT_UNCERTAINTY_ELLIPSE"
	SupportedGadShapes_POLYGON                    SupportedGadShapes = "POLYGON"
	SupportedGadShapes_POINT_ALTITUDE             SupportedGadShapes = "POINT_ALTITUDE"
	SupportedGadShapes_POINT_ALTITUDE_UNCERTAINTY SupportedGadShapes = "POINT_ALTITUDE_UNCERTAINTY"
	SupportedGadShapes_ELLIPSOID_ARC              SupportedGadShapes = "ELLIPSOID_ARC"
)
