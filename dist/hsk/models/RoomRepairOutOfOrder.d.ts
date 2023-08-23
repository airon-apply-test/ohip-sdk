/**
 * OPERA Cloud Housekeeping Service API
 * APIs to cater for Housekeeping functionality in OPERA Cloud. <br /><br />Housekeeping enables you to schedule daily room cleaning, maintenance, and housekeeping staff activities. It provides information on room status, out of order/out of service rooms, and forecasting.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { InstanceLink } from './InstanceLink';
import type { RoomRepairOutOfOrderCriteria } from './RoomRepairOutOfOrderCriteria';
import type { WarningType } from './WarningType';
/**
 *
 * @export
 * @interface RoomRepairOutOfOrder
 */
export interface RoomRepairOutOfOrder {
    /**
     *
     * @type {RoomRepairOutOfOrderCriteria}
     * @memberof RoomRepairOutOfOrder
     */
    criteria?: RoomRepairOutOfOrderCriteria;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof RoomRepairOutOfOrder
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof RoomRepairOutOfOrder
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the RoomRepairOutOfOrder interface.
 */
export declare function instanceOfRoomRepairOutOfOrder(value: object): boolean;
export declare function RoomRepairOutOfOrderFromJSON(json: any): RoomRepairOutOfOrder;
export declare function RoomRepairOutOfOrderFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoomRepairOutOfOrder;
export declare function RoomRepairOutOfOrderToJSON(value?: RoomRepairOutOfOrder | null): any;
