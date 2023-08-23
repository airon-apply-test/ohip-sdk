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
import type { LinkedAllowanceReservationsInfoType } from './LinkedAllowanceReservationsInfoType';
import type { PackagePostingType } from './PackagePostingType';
import type { TrxInfoType } from './TrxInfoType';
import type { WarningType } from './WarningType';
/**
 * Response to fetch reservation package options.
 * @export
 * @interface ReservationPackageOptions
 */
export interface ReservationPackageOptions {
    /**
     * List of package postings.
     * @type {Array<PackagePostingType>}
     * @memberof ReservationPackageOptions
     */
    packagePostings?: Array<PackagePostingType>;
    /**
     * List of Transaction codes info.
     * @type {Array<TrxInfoType>}
     * @memberof ReservationPackageOptions
     */
    trxCodesInfo?: Array<TrxInfoType>;
    /**
     *
     * @type {LinkedAllowanceReservationsInfoType}
     * @memberof ReservationPackageOptions
     */
    linkedAllowanceReservationsInfo?: LinkedAllowanceReservationsInfoType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof ReservationPackageOptions
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof ReservationPackageOptions
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the ReservationPackageOptions interface.
 */
export declare function instanceOfReservationPackageOptions(value: object): boolean;
export declare function ReservationPackageOptionsFromJSON(json: any): ReservationPackageOptions;
export declare function ReservationPackageOptionsFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReservationPackageOptions;
export declare function ReservationPackageOptionsToJSON(value?: ReservationPackageOptions | null): any;
