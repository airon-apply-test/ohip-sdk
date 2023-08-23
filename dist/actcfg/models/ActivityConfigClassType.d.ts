/**
 * OPERA Cloud Activity Management API
 * APIs to cater for Activity Configuration functionality in OPERA Cloud. In this module you can retrieve, create, update Activity configuration codes, for example create a new Activity Type.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * A collection of supported list of activity classes.
 * @export
 */
export declare const ActivityConfigClassType: {
    readonly Appointment: "Appointment";
    readonly Todo: "Todo";
};
export type ActivityConfigClassType = typeof ActivityConfigClassType[keyof typeof ActivityConfigClassType];
export declare function ActivityConfigClassTypeFromJSON(json: any): ActivityConfigClassType;
export declare function ActivityConfigClassTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ActivityConfigClassType;
export declare function ActivityConfigClassTypeToJSON(value?: ActivityConfigClassType | null): any;
