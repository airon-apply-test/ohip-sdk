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
import type { DepositFolioValidateCriteriaType } from './DepositFolioValidateCriteriaType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Request to validate a Deposit Folio for a reservation.
 * @export
 * @interface DepositFolioValidateCriteria
 */
export interface DepositFolioValidateCriteria {
    /**
     *
     * @type {DepositFolioValidateCriteriaType}
     * @memberof DepositFolioValidateCriteria
     */
    criteria?: DepositFolioValidateCriteriaType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof DepositFolioValidateCriteria
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof DepositFolioValidateCriteria
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the DepositFolioValidateCriteria interface.
 */
export declare function instanceOfDepositFolioValidateCriteria(value: object): boolean;
export declare function DepositFolioValidateCriteriaFromJSON(json: any): DepositFolioValidateCriteria;
export declare function DepositFolioValidateCriteriaFromJSONTyped(json: any, ignoreDiscriminator: boolean): DepositFolioValidateCriteria;
export declare function DepositFolioValidateCriteriaToJSON(value?: DepositFolioValidateCriteria | null): any;
