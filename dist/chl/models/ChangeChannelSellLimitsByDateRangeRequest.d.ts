/**
 * OPERA Cloud Channel Configuration API
 * APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ChannelSellLimitScheduleType } from './ChannelSellLimitScheduleType';
import type { WarningType } from './WarningType';
/**
 *
 * @export
 * @interface ChangeChannelSellLimitsByDateRangeRequest
 */
export interface ChangeChannelSellLimitsByDateRangeRequest {
    /**
     * Details about a sell limit schedule for a channel or channel room type.
     * @type {Array<ChannelSellLimitScheduleType>}
     * @memberof ChangeChannelSellLimitsByDateRangeRequest
     */
    sellLimits?: Array<ChannelSellLimitScheduleType>;
    /**
     * Flag to indicate whether any overlapping schedules should be automatically adjusted (split, truncated, etc.) as needed.
     * @type {boolean}
     * @memberof ChangeChannelSellLimitsByDateRangeRequest
     */
    adjustOverlappingSchedules?: boolean;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof ChangeChannelSellLimitsByDateRangeRequest
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the ChangeChannelSellLimitsByDateRangeRequest interface.
 */
export declare function instanceOfChangeChannelSellLimitsByDateRangeRequest(value: object): boolean;
export declare function ChangeChannelSellLimitsByDateRangeRequestFromJSON(json: any): ChangeChannelSellLimitsByDateRangeRequest;
export declare function ChangeChannelSellLimitsByDateRangeRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChangeChannelSellLimitsByDateRangeRequest;
export declare function ChangeChannelSellLimitsByDateRangeRequestToJSON(value?: ChangeChannelSellLimitsByDateRangeRequest | null): any;
