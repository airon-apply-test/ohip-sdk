/**
 * OPERA Cloud Block Configuration API
 * APIs for Block configuration, such as creating, updating, fetching and removing codes related to blocks. <br />< This might include fetching the block cancellation reasons, or creating new block refused reasons.  Wash schedules can be create, or new reservation methods could be added for a property.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { WashByRoomType } from './WashByRoomType';
/**
 * Allows for a choice to populate the detail information based on the choice between 'Wash by number of rooms' and 'Wash by %'.
 * @export
 * @interface WashCodeDetailType
 */
export interface WashCodeDetailType {
    /**
     *
     * @type {WashByRoomType}
     * @memberof WashCodeDetailType
     */
    washByRoom?: WashByRoomType;
    /**
     *
     * @type {number}
     * @memberof WashCodeDetailType
     */
    washByPercent?: number;
    /**
     * Number of days prior to arrival of a block or block's cutoff day.
     * @type {number}
     * @memberof WashCodeDetailType
     */
    daysPriorToArrival?: number;
}
/**
 * Check if a given object implements the WashCodeDetailType interface.
 */
export declare function instanceOfWashCodeDetailType(value: object): boolean;
export declare function WashCodeDetailTypeFromJSON(json: any): WashCodeDetailType;
export declare function WashCodeDetailTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): WashCodeDetailType;
export declare function WashCodeDetailTypeToJSON(value?: WashCodeDetailType | null): any;
