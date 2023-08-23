/**
 * OPERA Cloud Xchange Interface OXI API
 * APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 23.0.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ConversionCodeType } from './ConversionCodeType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 *
 * @export
 * @interface FetchConversionCodes
 */
export interface FetchConversionCodes {
    /**
     * List of Conversion Codes.
     * @type {Array<ConversionCodeType>}
     * @memberof FetchConversionCodes
     */
    conversionCodes?: Array<ConversionCodeType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof FetchConversionCodes
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof FetchConversionCodes
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the FetchConversionCodes interface.
 */
export declare function instanceOfFetchConversionCodes(value: object): boolean;
export declare function FetchConversionCodesFromJSON(json: any): FetchConversionCodes;
export declare function FetchConversionCodesFromJSONTyped(json: any, ignoreDiscriminator: boolean): FetchConversionCodes;
export declare function FetchConversionCodesToJSON(value?: FetchConversionCodes | null): any;
