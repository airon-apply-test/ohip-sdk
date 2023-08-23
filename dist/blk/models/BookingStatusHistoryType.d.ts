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
import type { CodeDescriptionType } from './CodeDescriptionType';
/**
 *
 * @export
 * @interface BookingStatusHistoryType
 */
export interface BookingStatusHistoryType {
    /**
     * Number indicating the sequence of status change.
     * @type {number}
     * @memberof BookingStatusHistoryType
     */
    sequence?: number;
    /**
     *
     * @type {CodeDescriptionType}
     * @memberof BookingStatusHistoryType
     */
    bookingStatus?: CodeDescriptionType;
    /**
     * Date and time of the status change.
     * @type {Date}
     * @memberof BookingStatusHistoryType
     */
    modifyDateTime?: Date;
    /**
     * User who modified the status.
     * @type {string}
     * @memberof BookingStatusHistoryType
     */
    modifierId?: string;
}
/**
 * Check if a given object implements the BookingStatusHistoryType interface.
 */
export declare function instanceOfBookingStatusHistoryType(value: object): boolean;
export declare function BookingStatusHistoryTypeFromJSON(json: any): BookingStatusHistoryType;
export declare function BookingStatusHistoryTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BookingStatusHistoryType;
export declare function BookingStatusHistoryTypeToJSON(value?: BookingStatusHistoryType | null): any;
