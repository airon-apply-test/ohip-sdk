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
import type { HotelInfoTypePropertyControlsApplicationMode } from './HotelInfoTypePropertyControlsApplicationMode';
import type { HotelInfoTypePropertyControlsCateringCurrencyFormatting } from './HotelInfoTypePropertyControlsCateringCurrencyFormatting';
import type { HotelInfoTypePropertyControlsCurrencyFormatting } from './HotelInfoTypePropertyControlsCurrencyFormatting';
import type { HotelInfoTypePropertyControlsDateTimeFormatting } from './HotelInfoTypePropertyControlsDateTimeFormatting';
import type { HotelInfoTypePropertyControlsSellControls } from './HotelInfoTypePropertyControlsSellControls';
/**
 * Property controls information configuration of the hotel.
 * @export
 * @interface HotelInfoTypePropertyControls
 */
export interface HotelInfoTypePropertyControls {
    /**
     *
     * @type {HotelInfoTypePropertyControlsSellControls}
     * @memberof HotelInfoTypePropertyControls
     */
    sellControls?: HotelInfoTypePropertyControlsSellControls;
    /**
     *
     * @type {HotelInfoTypePropertyControlsCurrencyFormatting}
     * @memberof HotelInfoTypePropertyControls
     */
    currencyFormatting?: HotelInfoTypePropertyControlsCurrencyFormatting;
    /**
     *
     * @type {HotelInfoTypePropertyControlsCateringCurrencyFormatting}
     * @memberof HotelInfoTypePropertyControls
     */
    cateringCurrencyFormatting?: HotelInfoTypePropertyControlsCateringCurrencyFormatting;
    /**
     *
     * @type {HotelInfoTypePropertyControlsDateTimeFormatting}
     * @memberof HotelInfoTypePropertyControls
     */
    dateTimeFormatting?: HotelInfoTypePropertyControlsDateTimeFormatting;
    /**
     *
     * @type {HotelInfoTypePropertyControlsApplicationMode}
     * @memberof HotelInfoTypePropertyControls
     */
    applicationMode?: HotelInfoTypePropertyControlsApplicationMode;
}
/**
 * Check if a given object implements the HotelInfoTypePropertyControls interface.
 */
export declare function instanceOfHotelInfoTypePropertyControls(value: object): boolean;
export declare function HotelInfoTypePropertyControlsFromJSON(json: any): HotelInfoTypePropertyControls;
export declare function HotelInfoTypePropertyControlsFromJSONTyped(json: any, ignoreDiscriminator: boolean): HotelInfoTypePropertyControls;
export declare function HotelInfoTypePropertyControlsToJSON(value?: HotelInfoTypePropertyControls | null): any;
