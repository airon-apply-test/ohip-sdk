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
import type { ECertificateType } from './ECertificateType';
import type { InstanceLink } from './InstanceLink';
/**
 * Response object for fetch ECertificates. This object contains collection of ECertificates,Success,Warnings and Errors related to this operation.
 * @export
 * @interface ECertificates
 */
export interface ECertificates {
    /**
     * List of e-certificates for the profile.
     * @type {Array<ECertificateType>}
     * @memberof ECertificates
     */
    eCertificatesDetail?: Array<ECertificateType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof ECertificates
     */
    links?: Array<InstanceLink>;
}
/**
 * Check if a given object implements the ECertificates interface.
 */
export declare function instanceOfECertificates(value: object): boolean;
export declare function ECertificatesFromJSON(json: any): ECertificates;
export declare function ECertificatesFromJSONTyped(json: any, ignoreDiscriminator: boolean): ECertificates;
export declare function ECertificatesToJSON(value?: ECertificates | null): any;
