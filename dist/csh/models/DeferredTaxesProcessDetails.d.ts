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
 * Response of submitting post deferred taxes for a reservation.
 * @export
 * @interface DeferredTaxesProcessDetails
 */
export interface DeferredTaxesProcessDetails {
    /**
     * Used for Character Strings, length 0 to 40.
     * @type {string}
     * @memberof DeferredTaxesProcessDetails
     */
    processId?: string;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof DeferredTaxesProcessDetails
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof DeferredTaxesProcessDetails
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the DeferredTaxesProcessDetails interface.
 */
export declare function instanceOfDeferredTaxesProcessDetails(value: object): boolean;
export declare function DeferredTaxesProcessDetailsFromJSON(json: any): DeferredTaxesProcessDetails;
export declare function DeferredTaxesProcessDetailsFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeferredTaxesProcessDetails;
export declare function DeferredTaxesProcessDetailsToJSON(value?: DeferredTaxesProcessDetails | null): any;
