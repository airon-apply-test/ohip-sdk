/**
 * OPERA Cloud Room Configuration API
 * APIs to cater for room configuration, such as configuring room types, room Classes, creating new room features, or updating housekeeping room maintenance reasons.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Indicates room type is a suite.
 * @export
 * @interface RoomTypeTypeSuite
 */
export interface RoomTypeTypeSuite {
    /**
     * Property Value
     * @type {boolean}
     * @memberof RoomTypeTypeSuite
     */
    value?: boolean;
    /**
     * Indicates if room type Room Components may be changed.
     * @type {boolean}
     * @memberof RoomTypeTypeSuite
     */
    editable?: boolean;
}
/**
 * Check if a given object implements the RoomTypeTypeSuite interface.
 */
export declare function instanceOfRoomTypeTypeSuite(value: object): boolean;
export declare function RoomTypeTypeSuiteFromJSON(json: any): RoomTypeTypeSuite;
export declare function RoomTypeTypeSuiteFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoomTypeTypeSuite;
export declare function RoomTypeTypeSuiteToJSON(value?: RoomTypeTypeSuite | null): any;
