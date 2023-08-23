/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { EmailInfoType } from './EmailInfoType';
/**
 * List of email address for the customer.
 * @export
 * @interface ProfileTypeEmails
 */
export interface ProfileTypeEmails {
    /**
     * Collection of Detailed information on an eMail address for the customer.
     * @type {Array<EmailInfoType>}
     * @memberof ProfileTypeEmails
     */
    emailInfo?: Array<EmailInfoType>;
    /**
     * Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
     * @type {boolean}
     * @memberof ProfileTypeEmails
     */
    hasMore?: boolean;
    /**
     * Total number of rows queried
     * @type {number}
     * @memberof ProfileTypeEmails
     */
    totalResults?: number;
    /**
     * Total number of rows returned
     * @type {number}
     * @memberof ProfileTypeEmails
     */
    count?: number;
}
/**
 * Check if a given object implements the ProfileTypeEmails interface.
 */
export declare function instanceOfProfileTypeEmails(value: object): boolean;
export declare function ProfileTypeEmailsFromJSON(json: any): ProfileTypeEmails;
export declare function ProfileTypeEmailsFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileTypeEmails;
export declare function ProfileTypeEmailsToJSON(value?: ProfileTypeEmails | null): any;
