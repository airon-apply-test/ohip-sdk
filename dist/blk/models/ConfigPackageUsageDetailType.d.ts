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
import type { ProductSourceType } from './ProductSourceType';
/**
 * A Config Package Usage Detail type.
 * @export
 * @interface ConfigPackageUsageDetailType
 */
export interface ConfigPackageUsageDetailType {
    /**
     *
     * @type {ProductSourceType}
     * @memberof ConfigPackageUsageDetailType
     */
    source?: ProductSourceType;
    /**
     * Indicates if the package is used in, reserved or prospect reservations.
     * @type {boolean}
     * @memberof ConfigPackageUsageDetailType
     */
    usedInReservations?: boolean;
    /**
     * Indicates if the package is used in any rate code.
     * @type {boolean}
     * @memberof ConfigPackageUsageDetailType
     */
    usedInRates?: boolean;
    /**
     * Indicates if any checked in reservations are using this product.
     * @type {boolean}
     * @memberof ConfigPackageUsageDetailType
     */
    usedInHouseReservations?: boolean;
}
/**
 * Check if a given object implements the ConfigPackageUsageDetailType interface.
 */
export declare function instanceOfConfigPackageUsageDetailType(value: object): boolean;
export declare function ConfigPackageUsageDetailTypeFromJSON(json: any): ConfigPackageUsageDetailType;
export declare function ConfigPackageUsageDetailTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ConfigPackageUsageDetailType;
export declare function ConfigPackageUsageDetailTypeToJSON(value?: ConfigPackageUsageDetailType | null): any;
