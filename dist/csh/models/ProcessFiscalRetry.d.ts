/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { FiscalRetryType } from './FiscalRetryType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * The request object to process fiscal retry functionality.
 * @export
 * @interface ProcessFiscalRetry
 */
export interface ProcessFiscalRetry {
    /**
     *
     * @type {FiscalRetryType}
     * @memberof ProcessFiscalRetry
     */
    criteria?: FiscalRetryType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof ProcessFiscalRetry
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof ProcessFiscalRetry
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the ProcessFiscalRetry interface.
 */
export declare function instanceOfProcessFiscalRetry(value: object): boolean;
export declare function ProcessFiscalRetryFromJSON(json: any): ProcessFiscalRetry;
export declare function ProcessFiscalRetryFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProcessFiscalRetry;
export declare function ProcessFiscalRetryToJSON(value?: ProcessFiscalRetry | null): any;
