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
import type { BlockId } from './BlockId';
import type { BorrowInventoryType } from './BorrowInventoryType';
import type { InstanceLink } from './InstanceLink';
import type { UniqueIDType } from './UniqueIDType';
import type { WarningType } from './WarningType';
/**
 *
 * @export
 * @interface PutInventoryRequest
 */
export interface PutInventoryRequest {
    /**
     * Used for codes in the OPERA Code tables. Possible values of this pattern are 1, 101, 101.EQP, or 101.EQP.X.
     * @type {string}
     * @memberof PutInventoryRequest
     */
    hotelId?: string;
    /**
     *
     * @type {BlockId}
     * @memberof PutInventoryRequest
     */
    blockId?: BlockId;
    /**
     *
     * @type {UniqueIDType}
     * @memberof PutInventoryRequest
     */
    existingReservationId?: UniqueIDType;
    /**
     * Used for codes in the OPERA Code tables. Possible values of this pattern are 1, 101, 101.EQP, or 101.EQP.X.
     * @type {string}
     * @memberof PutInventoryRequest
     */
    roomType?: string;
    /**
     * The number of adults that are expected to stay in a room.
     * @type {number}
     * @memberof PutInventoryRequest
     */
    adultCount?: number;
    /**
     * Indicates whether to overbook the Sales Allowance in case there are no rooms left at the Generic Sales Allowance level.
     * @type {boolean}
     * @memberof PutInventoryRequest
     */
    overbookSalesAllowance?: boolean;
    /**
     * The date and number of rooms to borrow from either room type or House.
     * @type {Array<BorrowInventoryType>}
     * @memberof PutInventoryRequest
     */
    borrowInventoryList?: Array<BorrowInventoryType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof PutInventoryRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PutInventoryRequest
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the PutInventoryRequest interface.
 */
export declare function instanceOfPutInventoryRequest(value: object): boolean;
export declare function PutInventoryRequestFromJSON(json: any): PutInventoryRequest;
export declare function PutInventoryRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PutInventoryRequest;
export declare function PutInventoryRequestToJSON(value?: PutInventoryRequest | null): any;
