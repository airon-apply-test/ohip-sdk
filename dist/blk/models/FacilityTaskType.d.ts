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
import type { FacilityCodeType } from './FacilityCodeType';
import type { HousekeepingTaskCodeType } from './HousekeepingTaskCodeType';
/**
 * Information regarding facility task on a reservation.
 * @export
 * @interface FacilityTaskType
 */
export interface FacilityTaskType {
    /**
     *
     * @type {HousekeepingTaskCodeType}
     * @memberof FacilityTaskType
     */
    task?: HousekeepingTaskCodeType;
    /**
     * List of the facility codes.
     * @type {Array<FacilityCodeType>}
     * @memberof FacilityTaskType
     */
    supplies?: Array<FacilityCodeType>;
    /**
     * The Date on which the task is applicable.
     * @type {Date}
     * @memberof FacilityTaskType
     */
    date?: Date;
}
/**
 * Check if a given object implements the FacilityTaskType interface.
 */
export declare function instanceOfFacilityTaskType(value: object): boolean;
export declare function FacilityTaskTypeFromJSON(json: any): FacilityTaskType;
export declare function FacilityTaskTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): FacilityTaskType;
export declare function FacilityTaskTypeToJSON(value?: FacilityTaskType | null): any;
