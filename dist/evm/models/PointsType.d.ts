/**
 * OPERA Cloud Sales Event Management API
 * APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * The Redeem Points contains the number of points required to book a functionSpaceDetails type.
 * @export
 * @interface PointsType
 */
export interface PointsType {
    /**
     * Awards type for the award type rate code.
     * @type {string}
     * @memberof PointsType
     */
    awardsType?: string;
    /**
     * No of points required to book this Room Stay.
     * @type {number}
     * @memberof PointsType
     */
    points?: number;
}
/**
 * Check if a given object implements the PointsType interface.
 */
export declare function instanceOfPointsType(value: object): boolean;
export declare function PointsTypeFromJSON(json: any): PointsType;
export declare function PointsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PointsType;
export declare function PointsTypeToJSON(value?: PointsType | null): any;
