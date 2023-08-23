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
import type { ConversionCodeStatusType } from './ConversionCodeStatusType';
import type { InstanceLink } from './InstanceLink';
import type { IntegrationSystemType } from './IntegrationSystemType';
import type { WarningType } from './WarningType';
/**
 *
 * @export
 * @interface ConversionCodesStatusToBeChanged
 */
export interface ConversionCodesStatusToBeChanged {
    /**
     *
     * @type {IntegrationSystemType}
     * @memberof ConversionCodesStatusToBeChanged
     */
    integrationSystem?: IntegrationSystemType;
    /**
     * Conversion Code and status information.
     * @type {Array<ConversionCodeStatusType>}
     * @memberof ConversionCodesStatusToBeChanged
     */
    conversionCodesStatus?: Array<ConversionCodeStatusType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof ConversionCodesStatusToBeChanged
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof ConversionCodesStatusToBeChanged
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the ConversionCodesStatusToBeChanged interface.
 */
export declare function instanceOfConversionCodesStatusToBeChanged(value: object): boolean;
export declare function ConversionCodesStatusToBeChangedFromJSON(json: any): ConversionCodesStatusToBeChanged;
export declare function ConversionCodesStatusToBeChangedFromJSONTyped(json: any, ignoreDiscriminator: boolean): ConversionCodesStatusToBeChanged;
export declare function ConversionCodesStatusToBeChangedToJSON(value?: ConversionCodesStatusToBeChanged | null): any;
