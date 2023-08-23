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
/**
 *
 * @export
 * @interface ChannelSellLimitType
 */
export interface ChannelSellLimitType {
    /**
     * Number of rooms that can be sold for this particular sell limit date.
     * @type {number}
     * @memberof ChannelSellLimitType
     */
    numberOfRooms?: number;
    /**
     * Channel for which this sell limit is applicable.
     * @type {string}
     * @memberof ChannelSellLimitType
     */
    bookingChannelCode?: string;
    /**
     * Date on which this sell limit is applicable.
     * @type {Date}
     * @memberof ChannelSellLimitType
     */
    date?: Date;
}
/**
 * Check if a given object implements the ChannelSellLimitType interface.
 */
export declare function instanceOfChannelSellLimitType(value: object): boolean;
export declare function ChannelSellLimitTypeFromJSON(json: any): ChannelSellLimitType;
export declare function ChannelSellLimitTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChannelSellLimitType;
export declare function ChannelSellLimitTypeToJSON(value?: ChannelSellLimitType | null): any;
