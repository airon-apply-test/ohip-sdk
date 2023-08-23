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
import type { TotalPricingElementType } from './TotalPricingElementType';
/**
 * Total Pricing Elements to copy from single source property to multi target properties
 * @export
 * @interface CopyTotalPricingElementsType
 */
export interface CopyTotalPricingElementsType {
    /**
     * List of Total Pricing Element Type
     * @type {Array<TotalPricingElementType>}
     * @memberof CopyTotalPricingElementsType
     */
    totalPricingElements?: Array<TotalPricingElementType>;
    /**
     *
     * @type {Array<string>}
     * @memberof CopyTotalPricingElementsType
     */
    targetHotels?: Array<string>;
}
/**
 * Check if a given object implements the CopyTotalPricingElementsType interface.
 */
export declare function instanceOfCopyTotalPricingElementsType(value: object): boolean;
export declare function CopyTotalPricingElementsTypeFromJSON(json: any): CopyTotalPricingElementsType;
export declare function CopyTotalPricingElementsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CopyTotalPricingElementsType;
export declare function CopyTotalPricingElementsTypeToJSON(value?: CopyTotalPricingElementsType | null): any;
