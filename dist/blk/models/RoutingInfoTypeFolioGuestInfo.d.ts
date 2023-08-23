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
import type { UniqueIDType } from './UniqueIDType';
/**
 * Guest details
 * @export
 * @interface RoutingInfoTypeFolioGuestInfo
 */
export interface RoutingInfoTypeFolioGuestInfo {
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof RoutingInfoTypeFolioGuestInfo
     */
    profileIdList?: Array<UniqueIDType>;
}
/**
 * Check if a given object implements the RoutingInfoTypeFolioGuestInfo interface.
 */
export declare function instanceOfRoutingInfoTypeFolioGuestInfo(value: object): boolean;
export declare function RoutingInfoTypeFolioGuestInfoFromJSON(json: any): RoutingInfoTypeFolioGuestInfo;
export declare function RoutingInfoTypeFolioGuestInfoFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoutingInfoTypeFolioGuestInfo;
export declare function RoutingInfoTypeFolioGuestInfoToJSON(value?: RoutingInfoTypeFolioGuestInfo | null): any;
