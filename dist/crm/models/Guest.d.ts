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
import type { ExternalReferenceType } from './ExternalReferenceType';
import type { GuestProfileType } from './GuestProfileType';
import type { InstanceLink } from './InstanceLink';
import type { UniqueIDType } from './UniqueIDType';
import type { WarningType } from './WarningType';
/**
 * Request object for creation of guest/contact/employee profile. This object contains profile details with unique identifiers of a profile. The standard optional Opera Context element is also included.
 * @export
 * @interface Guest
 */
export interface Guest {
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof Guest
     */
    guestIdList?: Array<UniqueIDType>;
    /**
     * This type contains unique information of external reference.
     * @type {Array<ExternalReferenceType>}
     * @memberof Guest
     */
    externalReferences?: Array<ExternalReferenceType>;
    /**
     *
     * @type {GuestProfileType}
     * @memberof Guest
     */
    guestDetails?: GuestProfileType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof Guest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof Guest
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the Guest interface.
 */
export declare function instanceOfGuest(value: object): boolean;
export declare function GuestFromJSON(json: any): Guest;
export declare function GuestFromJSONTyped(json: any, ignoreDiscriminator: boolean): Guest;
export declare function GuestToJSON(value?: Guest | null): any;
