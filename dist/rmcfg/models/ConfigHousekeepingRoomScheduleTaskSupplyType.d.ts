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
import type { CodeDescriptionType } from './CodeDescriptionType';
/**
 * Facility Housekeeping Code, its description and quantity.
 * @export
 * @interface ConfigHousekeepingRoomScheduleTaskSupplyType
 */
export interface ConfigHousekeepingRoomScheduleTaskSupplyType {
    /**
     * Facility Code.
     * @type {string}
     * @memberof ConfigHousekeepingRoomScheduleTaskSupplyType
     */
    description?: string;
    /**
     * Signifies the quantity.
     * @type {number}
     * @memberof ConfigHousekeepingRoomScheduleTaskSupplyType
     */
    quantity?: number;
    /**
     * Facility code value.
     * @type {string}
     * @memberof ConfigHousekeepingRoomScheduleTaskSupplyType
     */
    code?: string;
    /**
     * Used to store the housekeeping room schedule sequence.
     * @type {number}
     * @memberof ConfigHousekeepingRoomScheduleTaskSupplyType
     */
    sequence?: number;
    /**
     *
     * @type {CodeDescriptionType}
     * @memberof ConfigHousekeepingRoomScheduleTaskSupplyType
     */
    roomType?: CodeDescriptionType;
    /**
     * Specifies the hotel code that the room type belongs to.
     * @type {string}
     * @memberof ConfigHousekeepingRoomScheduleTaskSupplyType
     */
    hotelId?: string;
    /**
     * Specifies the housekeeping room schedule code.
     * @type {string}
     * @memberof ConfigHousekeepingRoomScheduleTaskSupplyType
     */
    housekeepingRoomScheduleCode?: string;
}
/**
 * Check if a given object implements the ConfigHousekeepingRoomScheduleTaskSupplyType interface.
 */
export declare function instanceOfConfigHousekeepingRoomScheduleTaskSupplyType(value: object): boolean;
export declare function ConfigHousekeepingRoomScheduleTaskSupplyTypeFromJSON(json: any): ConfigHousekeepingRoomScheduleTaskSupplyType;
export declare function ConfigHousekeepingRoomScheduleTaskSupplyTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ConfigHousekeepingRoomScheduleTaskSupplyType;
export declare function ConfigHousekeepingRoomScheduleTaskSupplyTypeToJSON(value?: ConfigHousekeepingRoomScheduleTaskSupplyType | null): any;
