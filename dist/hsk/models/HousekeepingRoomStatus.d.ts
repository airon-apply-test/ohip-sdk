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
import type { WarningType } from './WarningType';
/**
 * Existing Operations Responses will eventually be modified to be extended from this type.
 * @export
 * @interface HousekeepingRoomStatus
 */
export interface HousekeepingRoomStatus {
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof HousekeepingRoomStatus
     */
    warnings?: Array<WarningType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof HousekeepingRoomStatus
     */
    links?: Array<InstanceLink>;
}
/**
 * Check if a given object implements the HousekeepingRoomStatus interface.
 */
export declare function instanceOfHousekeepingRoomStatus(value: object): boolean;
export declare function HousekeepingRoomStatusFromJSON(json: any): HousekeepingRoomStatus;
export declare function HousekeepingRoomStatusFromJSONTyped(json: any, ignoreDiscriminator: boolean): HousekeepingRoomStatus;
export declare function HousekeepingRoomStatusToJSON(value?: HousekeepingRoomStatus | null): any;
