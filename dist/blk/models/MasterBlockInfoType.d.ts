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
import type { BlockClassificationType } from './BlockClassificationType';
import type { BookingStatusType } from './BookingStatusType';
import type { MasterSubBlockBaseInfoType } from './MasterSubBlockBaseInfoType';
import type { TimeSpanType } from './TimeSpanType';
import type { UniqueIDType } from './UniqueIDType';
/**
 * Basic information pertaining to the master block.
 * @export
 * @interface MasterBlockInfoType
 */
export interface MasterBlockInfoType {
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof MasterBlockInfoType
     */
    blockIdList?: Array<UniqueIDType>;
    /**
     * Master/Sub Block's Hotel Code.
     * @type {string}
     * @memberof MasterBlockInfoType
     */
    hotelId?: string;
    /**
     *
     * @type {BlockClassificationType}
     * @memberof MasterBlockInfoType
     */
    blockType?: BlockClassificationType;
    /**
     *
     * @type {TimeSpanType}
     * @memberof MasterBlockInfoType
     */
    timeSpan?: TimeSpanType;
    /**
     * Block description.
     * @type {string}
     * @memberof MasterBlockInfoType
     */
    blockName?: string;
    /**
     *
     * @type {BookingStatusType}
     * @memberof MasterBlockInfoType
     */
    blockStatus?: BookingStatusType;
    /**
     *
     * @type {BookingStatusType}
     * @memberof MasterBlockInfoType
     */
    cateringStatus?: BookingStatusType;
    /**
     * Pertain value for blocked rooms for a block.
     * @type {number}
     * @memberof MasterBlockInfoType
     */
    roomNights?: number;
    /**
     * Pertain value for reserved rooms for a block.
     * @type {number}
     * @memberof MasterBlockInfoType
     */
    roomNightsPickedup?: number;
    /**
     * Indicates to check if Block Dates to be in sync.
     * @type {boolean}
     * @memberof MasterBlockInfoType
     */
    syncBlockDates?: boolean;
    /**
     * Indicates whether other BlockDetails to be in Sync.
     * @type {boolean}
     * @memberof MasterBlockInfoType
     */
    syncOtherBlockDetails?: boolean;
    /**
     * Indicates if Block Status Details to be in sync.
     * @type {boolean}
     * @memberof MasterBlockInfoType
     */
    syncBlockStatusDetails?: boolean;
    /**
     * Indicates if packages to be in Sync.
     * @type {boolean}
     * @memberof MasterBlockInfoType
     */
    syncPackages?: boolean;
    /**
     * Basic information pertaining to the sub block.
     * @type {Array<MasterSubBlockBaseInfoType>}
     * @memberof MasterBlockInfoType
     */
    subBlockInfo?: Array<MasterSubBlockBaseInfoType>;
    /**
     * Pertains valid hotel code list for logged in user against Master/Sub header record.
     * @type {Array<Array<string>>}
     * @memberof MasterBlockInfoType
     */
    masterSubHotels?: Array<Array<string>>;
}
/**
 * Check if a given object implements the MasterBlockInfoType interface.
 */
export declare function instanceOfMasterBlockInfoType(value: object): boolean;
export declare function MasterBlockInfoTypeFromJSON(json: any): MasterBlockInfoType;
export declare function MasterBlockInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): MasterBlockInfoType;
export declare function MasterBlockInfoTypeToJSON(value?: MasterBlockInfoType | null): any;
