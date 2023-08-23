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
import type { TimeSpanType } from './TimeSpanType';
/**
 *
 * @export
 * @interface CreateSubBlockBaseInfoType
 */
export interface CreateSubBlockBaseInfoType {
    /**
     * Sub Block's Hotel Code.
     * @type {string}
     * @memberof CreateSubBlockBaseInfoType
     */
    hotelId?: string;
    /**
     * Block description.
     * @type {string}
     * @memberof CreateSubBlockBaseInfoType
     */
    blockName?: string;
    /**
     * Block code for the sub block.
     * @type {string}
     * @memberof CreateSubBlockBaseInfoType
     */
    blockCode?: string;
    /**
     * Origin information of the block PMS/ORS/SC/SFA
     * @type {string}
     * @memberof CreateSubBlockBaseInfoType
     */
    blockOrigin?: string;
    /**
     *
     * @type {TimeSpanType}
     * @memberof CreateSubBlockBaseInfoType
     */
    timeSpan?: TimeSpanType;
}
/**
 * Check if a given object implements the CreateSubBlockBaseInfoType interface.
 */
export declare function instanceOfCreateSubBlockBaseInfoType(value: object): boolean;
export declare function CreateSubBlockBaseInfoTypeFromJSON(json: any): CreateSubBlockBaseInfoType;
export declare function CreateSubBlockBaseInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateSubBlockBaseInfoType;
export declare function CreateSubBlockBaseInfoTypeToJSON(value?: CreateSubBlockBaseInfoType | null): any;
