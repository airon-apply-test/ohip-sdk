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
import type { UniqueIDType } from './UniqueIDType';
/**
 * Key information about the block for a reservation.
 * @export
 * @interface ReservationBlockType
 */
export interface ReservationBlockType {
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof ReservationBlockType
     */
    blockIdList?: Array<UniqueIDType>;
    /**
     * The Name of the block that is attached to the reservation.
     * @type {string}
     * @memberof ReservationBlockType
     */
    blockName?: string;
    /**
     * This is the HotelCode of the Block.
     * @type {string}
     * @memberof ReservationBlockType
     */
    hotelId?: string;
}
/**
 * Check if a given object implements the ReservationBlockType interface.
 */
export declare function instanceOfReservationBlockType(value: object): boolean;
export declare function ReservationBlockTypeFromJSON(json: any): ReservationBlockType;
export declare function ReservationBlockTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReservationBlockType;
export declare function ReservationBlockTypeToJSON(value?: ReservationBlockType | null): any;
