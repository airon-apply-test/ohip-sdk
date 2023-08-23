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
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Check if it will call fiscal or not.
 * @export
 * @interface FiscalCallStatus
 */
export interface FiscalCallStatus {
    /**
     * Flag which contains if its valid or not.
     * @type {boolean}
     * @memberof FiscalCallStatus
     */
    isValid?: boolean;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof FiscalCallStatus
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof FiscalCallStatus
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the FiscalCallStatus interface.
 */
export declare function instanceOfFiscalCallStatus(value: object): boolean;
export declare function FiscalCallStatusFromJSON(json: any): FiscalCallStatus;
export declare function FiscalCallStatusFromJSONTyped(json: any, ignoreDiscriminator: boolean): FiscalCallStatus;
export declare function FiscalCallStatusToJSON(value?: FiscalCallStatus | null): any;
