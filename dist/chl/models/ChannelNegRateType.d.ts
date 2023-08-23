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
import type { ChannelNegProfileType } from './ChannelNegProfileType';
/**
 *
 * @export
 * @interface ChannelNegRateType
 */
export interface ChannelNegRateType {
    /**
     *
     * @type {Array<ChannelNegProfileType>}
     * @memberof ChannelNegRateType
     */
    negProfile?: Array<ChannelNegProfileType>;
    /**
     * Booking Channel Code.
     * @type {string}
     * @memberof ChannelNegRateType
     */
    bookingChannelCode?: string;
    /**
     * Hotel Code.
     * @type {string}
     * @memberof ChannelNegRateType
     */
    hotelId?: string;
    /**
     * Channel Room Type.
     * @type {string}
     * @memberof ChannelNegRateType
     */
    channelRatePlanCode?: string;
}
/**
 * Check if a given object implements the ChannelNegRateType interface.
 */
export declare function instanceOfChannelNegRateType(value: object): boolean;
export declare function ChannelNegRateTypeFromJSON(json: any): ChannelNegRateType;
export declare function ChannelNegRateTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChannelNegRateType;
export declare function ChannelNegRateTypeToJSON(value?: ChannelNegRateType | null): any;
