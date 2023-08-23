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
import type { ErrorType } from './ErrorType';
import type { WarningType } from './WarningType';
/**
 * This provides a collection of all borrowable room types by date.
 * @export
 * @interface BorrowableInventoryType
 */
export interface BorrowableInventoryType {
    /**
     * Returning an empty element of this type indicates the this is a House type. This is used in conjunction with the Borrow operations for Blocks where rooms are to be borrowed from House.
     * @type {object}
     * @memberof BorrowableInventoryType
     */
    house?: object;
    /**
     * The room type that contains rooms that can be borrowed.
     * @type {string}
     * @memberof BorrowableInventoryType
     */
    roomType?: string;
    /**
     * The number of rooms that are available to be borrowed from the room type or house.
     * @type {number}
     * @memberof BorrowableInventoryType
     */
    availableRooms?: number;
    /**
     * The number of sell limit rooms that are available for the room type.
     * @type {number}
     * @memberof BorrowableInventoryType
     */
    availableSellLimit?: number;
    /**
     * An error that occurred during the processing of a message.
     * @type {Array<ErrorType>}
     * @memberof BorrowableInventoryType
     */
    errors?: Array<ErrorType>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof BorrowableInventoryType
     */
    warning?: Array<WarningType>;
}
/**
 * Check if a given object implements the BorrowableInventoryType interface.
 */
export declare function instanceOfBorrowableInventoryType(value: object): boolean;
export declare function BorrowableInventoryTypeFromJSON(json: any): BorrowableInventoryType;
export declare function BorrowableInventoryTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BorrowableInventoryType;
export declare function BorrowableInventoryTypeToJSON(value?: BorrowableInventoryType | null): any;
