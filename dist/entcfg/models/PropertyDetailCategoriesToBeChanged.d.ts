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
import type { InstanceLink } from './InstanceLink';
import type { PropertyDetailCategoryType } from './PropertyDetailCategoryType';
import type { WarningType } from './WarningType';
/**
 * Request object for changing Property Detail Categories.
 * @export
 * @interface PropertyDetailCategoriesToBeChanged
 */
export interface PropertyDetailCategoriesToBeChanged {
    /**
     * List of Property Detail Categories.
     * @type {Array<PropertyDetailCategoryType>}
     * @memberof PropertyDetailCategoriesToBeChanged
     */
    propertyDetailCategories?: Array<PropertyDetailCategoryType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof PropertyDetailCategoriesToBeChanged
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PropertyDetailCategoriesToBeChanged
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the PropertyDetailCategoriesToBeChanged interface.
 */
export declare function instanceOfPropertyDetailCategoriesToBeChanged(value: object): boolean;
export declare function PropertyDetailCategoriesToBeChangedFromJSON(json: any): PropertyDetailCategoriesToBeChanged;
export declare function PropertyDetailCategoriesToBeChangedFromJSONTyped(json: any, ignoreDiscriminator: boolean): PropertyDetailCategoriesToBeChanged;
export declare function PropertyDetailCategoriesToBeChangedToJSON(value?: PropertyDetailCategoriesToBeChanged | null): any;
