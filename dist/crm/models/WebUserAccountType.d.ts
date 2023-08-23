/**
 * OPERA Cloud Customer Relationship Management API
 * APIs to cater for Customer Relationship Management (profile) functionality in OPERA Cloud.  There are different types of profiles in OPERA Cloud, including Guest, Company, Travel Agent, Source, Group, and Contact profile types.  A profile can store and display a wide range of information about the guest, company, travel agent etc., depending upon the type of profile.  For example, a guest profile can store the guest name, address, contact information, details on billing, membership benefits, preferences and much more.  All profiles in OPERA when created are assigned a ProfileID.  This ID will be used throughout the CRM APIs.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { CodeDescriptionType } from './CodeDescriptionType';
import type { WebUserHistoryType } from './WebUserHistoryType';
/**
 * Web User Account of the guest.
 * @export
 * @interface WebUserAccountType
 */
export interface WebUserAccountType {
    /**
     *
     * @type {CodeDescriptionType}
     * @memberof WebUserAccountType
     */
    securityQuestion?: CodeDescriptionType;
    /**
     * Answer to the Security Question.
     * @type {string}
     * @memberof WebUserAccountType
     */
    securityAnswer?: string;
    /**
     * Comments.
     * @type {string}
     * @memberof WebUserAccountType
     */
    comments?: string;
    /**
     *
     * @type {WebUserHistoryType}
     * @memberof WebUserAccountType
     */
    history?: WebUserHistoryType;
    /**
     * New Login Name of the guest.
     * @type {string}
     * @memberof WebUserAccountType
     */
    newLoginName?: string;
    /**
     * Login Password.
     * @type {string}
     * @memberof WebUserAccountType
     */
    newPassword?: string;
    /**
     * A flag which determines if the password is auto generated.
     * @type {boolean}
     * @memberof WebUserAccountType
     */
    autoGeneratePassword?: boolean;
    /**
     * Login Name of the guest.
     * @type {string}
     * @memberof WebUserAccountType
     */
    loginName?: string;
    /**
     * Domain code.
     * @type {string}
     * @memberof WebUserAccountType
     */
    domainCode?: string;
    /**
     * Boolean flag that indicates whether the web account is locked or not.
     * @type {boolean}
     * @memberof WebUserAccountType
     */
    locked?: boolean;
    /**
     * Boolean flag that indicates whether the web account is inactive or not.
     * @type {boolean}
     * @memberof WebUserAccountType
     */
    inactive?: boolean;
}
/**
 * Check if a given object implements the WebUserAccountType interface.
 */
export declare function instanceOfWebUserAccountType(value: object): boolean;
export declare function WebUserAccountTypeFromJSON(json: any): WebUserAccountType;
export declare function WebUserAccountTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): WebUserAccountType;
export declare function WebUserAccountTypeToJSON(value?: WebUserAccountType | null): any;
