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
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Response object containing sell limit schedules set for the channel or channel room type.
 * @export
 * @interface ChannelSellLimits
 */
export interface ChannelSellLimits {
    /**
     * Details about a sell limit schedule for a channel or channel room type.
     * @type {Array<ChannelSellLimitScheduleType>}
     * @memberof ChannelSellLimits
     */
    sellLimits?: Array<ChannelSellLimitScheduleType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof ChannelSellLimits
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof ChannelSellLimits
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the ChannelSellLimits interface.
 */
export declare function instanceOfChannelSellLimits(value: object): boolean;
export declare function ChannelSellLimitsFromJSON(json: any): ChannelSellLimits;
export declare function ChannelSellLimitsFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChannelSellLimits;
export declare function ChannelSellLimitsToJSON(value?: ChannelSellLimits | null): any;
