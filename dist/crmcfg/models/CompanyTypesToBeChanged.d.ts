/**
 * OPERA Cloud CRM Configuration API
 * APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { CompanyTypeType } from './CompanyTypeType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Request object for changing Company Types.
 * @export
 * @interface CompanyTypesToBeChanged
 */
export interface CompanyTypesToBeChanged {
    /**
     * List of Company Types.
     * @type {Array<CompanyTypeType>}
     * @memberof CompanyTypesToBeChanged
     */
    companyTypes?: Array<CompanyTypeType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof CompanyTypesToBeChanged
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof CompanyTypesToBeChanged
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the CompanyTypesToBeChanged interface.
 */
export declare function instanceOfCompanyTypesToBeChanged(value: object): boolean;
export declare function CompanyTypesToBeChangedFromJSON(json: any): CompanyTypesToBeChanged;
export declare function CompanyTypesToBeChangedFromJSONTyped(json: any, ignoreDiscriminator: boolean): CompanyTypesToBeChanged;
export declare function CompanyTypesToBeChangedToJSON(value?: CompanyTypesToBeChanged | null): any;
