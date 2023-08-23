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
import type { SystemRateLevelType } from './SystemRateLevelType';
/**
 * channel rate levels to be created.
 * @export
 * @interface ChannelRateLevelsRateLevels
 */
export interface ChannelRateLevelsRateLevels {
    /**
     * Channel rate level details.
     * @type {Array<SystemRateLevelType>}
     * @memberof ChannelRateLevelsRateLevels
     */
    rateLevel?: Array<SystemRateLevelType>;
    /**
     * Used for codes in the OPERA Code tables. Possible values of this pattern are 1, 101, 101.EQP, or 101.EQP.X.
     * @type {string}
     * @memberof ChannelRateLevelsRateLevels
     */
    bookingChannelCode?: string;
}
/**
 * Check if a given object implements the ChannelRateLevelsRateLevels interface.
 */
export declare function instanceOfChannelRateLevelsRateLevels(value: object): boolean;
export declare function ChannelRateLevelsRateLevelsFromJSON(json: any): ChannelRateLevelsRateLevels;
export declare function ChannelRateLevelsRateLevelsFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChannelRateLevelsRateLevels;
export declare function ChannelRateLevelsRateLevelsToJSON(value?: ChannelRateLevelsRateLevels | null): any;
