/**
 * OPERA Cloud Reservation API
 * APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Guest Preference details for the profile.
 * @export
 * @interface PreferenceType
 */
export interface PreferenceType {
    /**
     * Preference value for display purposes.
     * @type {string}
     * @memberof PreferenceType
     */
    preferenceValue?: string;
    /**
     * Preference Description for display purposes.
     * @type {string}
     * @memberof PreferenceType
     */
    description?: string;
    /**
     * Whether this preference is property specific or not.
     * @type {boolean}
     * @memberof PreferenceType
     */
    global?: boolean;
    /**
     * Source of the preference.
     * @type {string}
     * @memberof PreferenceType
     */
    source?: string;
    /**
     * If specified preference belongs to the Hotels listed, otherwise it is a global preference.
     * @type {Array<string>}
     * @memberof PreferenceType
     */
    hotels?: Array<string>;
    /**
     *
     * @type {string}
     * @memberof PreferenceType
     */
    preferenceId?: string;
    /**
     * Specifies the count of preferences excluded for the attached reservation preference.
     * @type {number}
     * @memberof PreferenceType
     */
    excludedPreferencesCount?: number;
    /**
     * Specifies whether to copy the reservation preference to the profile or not.
     * @type {boolean}
     * @memberof PreferenceType
     */
    copyToProfile?: boolean;
}
/**
 * Check if a given object implements the PreferenceType interface.
 */
export declare function instanceOfPreferenceType(value: object): boolean;
export declare function PreferenceTypeFromJSON(json: any): PreferenceType;
export declare function PreferenceTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PreferenceType;
export declare function PreferenceTypeToJSON(value?: PreferenceType | null): any;
