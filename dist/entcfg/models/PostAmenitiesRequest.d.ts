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
import type { ConfigHotelAmenityType } from './ConfigHotelAmenityType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 *
 * @export
 * @interface PostAmenitiesRequest
 */
export interface PostAmenitiesRequest {
    /**
     * This type holds a collection of amenities at the property level.
     * @type {Array<ConfigHotelAmenityType>}
     * @memberof PostAmenitiesRequest
     */
    hotelAmenities?: Array<ConfigHotelAmenityType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof PostAmenitiesRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PostAmenitiesRequest
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the PostAmenitiesRequest interface.
 */
export declare function instanceOfPostAmenitiesRequest(value: object): boolean;
export declare function PostAmenitiesRequestFromJSON(json: any): PostAmenitiesRequest;
export declare function PostAmenitiesRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostAmenitiesRequest;
export declare function PostAmenitiesRequestToJSON(value?: PostAmenitiesRequest | null): any;
