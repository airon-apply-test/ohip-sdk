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
 * Request object for creating amenities at the property level.
 * @export
 * @interface AmenitiesCriteria
 */
export interface AmenitiesCriteria {
    /**
     * This type holds a collection of amenities at the property level.
     * @type {Array<ConfigHotelAmenityType>}
     * @memberof AmenitiesCriteria
     */
    hotelAmenities?: Array<ConfigHotelAmenityType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof AmenitiesCriteria
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof AmenitiesCriteria
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the AmenitiesCriteria interface.
 */
export declare function instanceOfAmenitiesCriteria(value: object): boolean;
export declare function AmenitiesCriteriaFromJSON(json: any): AmenitiesCriteria;
export declare function AmenitiesCriteriaFromJSONTyped(json: any, ignoreDiscriminator: boolean): AmenitiesCriteria;
export declare function AmenitiesCriteriaToJSON(value?: AmenitiesCriteria | null): any;
