/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { CommentInfoType } from './CommentInfoType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Request object for changing hotel Notes for hotels.
 * @export
 * @interface HotelNotesToBeChanged
 */
export interface HotelNotesToBeChanged {
    /**
     * List of Notes of the hotel.
     * @type {Array<CommentInfoType>}
     * @memberof HotelNotesToBeChanged
     */
    hotelNotes?: Array<CommentInfoType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof HotelNotesToBeChanged
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof HotelNotesToBeChanged
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the HotelNotesToBeChanged interface.
 */
export declare function instanceOfHotelNotesToBeChanged(value: object): boolean;
export declare function HotelNotesToBeChangedFromJSON(json: any): HotelNotesToBeChanged;
export declare function HotelNotesToBeChangedFromJSONTyped(json: any, ignoreDiscriminator: boolean): HotelNotesToBeChanged;
export declare function HotelNotesToBeChangedToJSON(value?: HotelNotesToBeChanged | null): any;
