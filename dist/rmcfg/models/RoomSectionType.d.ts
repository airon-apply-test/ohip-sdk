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
 *
 * @export
 * @interface RoomSectionType
 */
export interface RoomSectionType {
    /**
     *
     * @type {string}
     * @memberof RoomSectionType
     */
    daySectionCode?: string;
    /**
     *
     * @type {string}
     * @memberof RoomSectionType
     */
    eveningSectionCode?: string;
}
/**
 * Check if a given object implements the RoomSectionType interface.
 */
export declare function instanceOfRoomSectionType(value: object): boolean;
export declare function RoomSectionTypeFromJSON(json: any): RoomSectionType;
export declare function RoomSectionTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoomSectionType;
export declare function RoomSectionTypeToJSON(value?: RoomSectionType | null): any;
