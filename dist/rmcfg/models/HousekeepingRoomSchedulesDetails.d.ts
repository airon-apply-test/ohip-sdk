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
import type { ConfigHousekeepingRoomScheduleType } from './ConfigHousekeepingRoomScheduleType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Response object for fetching housekeeping tasks and housekeeping codes associated to a room type.
 * @export
 * @interface HousekeepingRoomSchedulesDetails
 */
export interface HousekeepingRoomSchedulesDetails {
    /**
     * This type holds a collection of housekeeping tasks attached to a room type.
     * @type {Array<ConfigHousekeepingRoomScheduleType>}
     * @memberof HousekeepingRoomSchedulesDetails
     */
    housekeepingRoomSchedules?: Array<ConfigHousekeepingRoomScheduleType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof HousekeepingRoomSchedulesDetails
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof HousekeepingRoomSchedulesDetails
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the HousekeepingRoomSchedulesDetails interface.
 */
export declare function instanceOfHousekeepingRoomSchedulesDetails(value: object): boolean;
export declare function HousekeepingRoomSchedulesDetailsFromJSON(json: any): HousekeepingRoomSchedulesDetails;
export declare function HousekeepingRoomSchedulesDetailsFromJSONTyped(json: any, ignoreDiscriminator: boolean): HousekeepingRoomSchedulesDetails;
export declare function HousekeepingRoomSchedulesDetailsToJSON(value?: HousekeepingRoomSchedulesDetails | null): any;
