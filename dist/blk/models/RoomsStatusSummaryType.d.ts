/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Contains Block's room status summary info.
 * @export
 * @interface RoomsStatusSummaryType
 */
export interface RoomsStatusSummaryType {
    /**
     *
     * @type {number}
     * @memberof RoomsStatusSummaryType
     */
    assigned?: number;
    /**
     *
     * @type {number}
     * @memberof RoomsStatusSummaryType
     */
    unassigned?: number;
    /**
     *
     * @type {number}
     * @memberof RoomsStatusSummaryType
     */
    inspected?: number;
    /**
     *
     * @type {number}
     * @memberof RoomsStatusSummaryType
     */
    dirty?: number;
    /**
     *
     * @type {number}
     * @memberof RoomsStatusSummaryType
     */
    clean?: number;
    /**
     *
     * @type {number}
     * @memberof RoomsStatusSummaryType
     */
    pickup?: number;
    /**
     *
     * @type {number}
     * @memberof RoomsStatusSummaryType
     */
    outOfService?: number;
}
/**
 * Check if a given object implements the RoomsStatusSummaryType interface.
 */
export declare function instanceOfRoomsStatusSummaryType(value: object): boolean;
export declare function RoomsStatusSummaryTypeFromJSON(json: any): RoomsStatusSummaryType;
export declare function RoomsStatusSummaryTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoomsStatusSummaryType;
export declare function RoomsStatusSummaryTypeToJSON(value?: RoomsStatusSummaryType | null): any;
