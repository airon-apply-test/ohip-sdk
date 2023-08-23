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
import type { DetailPostingType } from './DetailPostingType';
import type { FolioWindowType } from './FolioWindowType';
import type { InstanceLink } from './InstanceLink';
import type { TrxInfoType } from './TrxInfoType';
import type { WarningType } from './WarningType';
/**
 * Response for the request to create Supplemental Folio.
 * @export
 * @interface PostSupplementalFolioStatus
 */
export interface PostSupplementalFolioStatus {
    /**
     * Information regarding the new Supplemental Folio created.
     * @type {Array<FolioWindowType>}
     * @memberof PostSupplementalFolioStatus
     */
    folioWindow?: Array<FolioWindowType>;
    /**
     * Details of the transaction(Posting).
     * @type {Array<DetailPostingType>}
     * @memberof PostSupplementalFolioStatus
     */
    payments?: Array<DetailPostingType>;
    /**
     * List of Transaction codes info.
     * @type {Array<TrxInfoType>}
     * @memberof PostSupplementalFolioStatus
     */
    trxCodesInfo?: Array<TrxInfoType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof PostSupplementalFolioStatus
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PostSupplementalFolioStatus
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the PostSupplementalFolioStatus interface.
 */
export declare function instanceOfPostSupplementalFolioStatus(value: object): boolean;
export declare function PostSupplementalFolioStatusFromJSON(json: any): PostSupplementalFolioStatus;
export declare function PostSupplementalFolioStatusFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostSupplementalFolioStatus;
export declare function PostSupplementalFolioStatusToJSON(value?: PostSupplementalFolioStatus | null): any;
