/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ECertificateIssueSourceType } from './ECertificateIssueSourceType';
import type { ECertificateIssueType } from './ECertificateIssueType';
import type { ProfileId } from './ProfileId';
/**
 * E-Certificates issue API call details.
 * @export
 * @interface ECertificateGenerationDetailsType
 */
export interface ECertificateGenerationDetailsType {
    /**
     *
     * @type {ProfileId}
     * @memberof ECertificateGenerationDetailsType
     */
    profileId?: ProfileId;
    /**
     * User defined certificate code.
     * @type {string}
     * @memberof ECertificateGenerationDetailsType
     */
    certificateType?: string;
    /**
     *
     * @type {ECertificateIssueSourceType}
     * @memberof ECertificateGenerationDetailsType
     */
    source?: ECertificateIssueSourceType;
    /**
     *
     * @type {ECertificateIssueType}
     * @memberof ECertificateGenerationDetailsType
     */
    issueType?: ECertificateIssueType;
    /**
     * Hotel to which certificate is attached to.
     * @type {Array<string>}
     * @memberof ECertificateGenerationDetailsType
     */
    hotels?: Array<string>;
}
/**
 * Check if a given object implements the ECertificateGenerationDetailsType interface.
 */
export declare function instanceOfECertificateGenerationDetailsType(value: object): boolean;
export declare function ECertificateGenerationDetailsTypeFromJSON(json: any): ECertificateGenerationDetailsType;
export declare function ECertificateGenerationDetailsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ECertificateGenerationDetailsType;
export declare function ECertificateGenerationDetailsTypeToJSON(value?: ECertificateGenerationDetailsType | null): any;
