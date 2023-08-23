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
import type { PersonNameTypeType } from './PersonNameTypeType';
/**
 * This provides name information for a person.
 * @export
 * @interface PersonNameType
 */
export interface PersonNameType {
    /**
     * Salutation of honorific (e.g. Mr., Mrs., Ms., Miss, Dr.)
     * @type {string}
     * @memberof PersonNameType
     */
    namePrefix?: string;
    /**
     * Given name, first name or names.
     * @type {string}
     * @memberof PersonNameType
     */
    givenName?: string;
    /**
     * The middle name of the person name.
     * @type {string}
     * @memberof PersonNameType
     */
    middleName?: string;
    /**
     * Family name, last name. May also be used for full name if the sending system does not have the ability to separate a full name into its parts, e.g. the surname element may be used to pass the full name.
     * @type {string}
     * @memberof PersonNameType
     */
    surname?: string;
    /**
     * Hold various name suffixes and letters (e.g. Jr., Sr., III, Ret., Esq.)
     * @type {string}
     * @memberof PersonNameType
     */
    nameSuffix?: string;
    /**
     * Degree or honors (e.g., Ph.D., M.D.)
     * @type {string}
     * @memberof PersonNameType
     */
    nameTitle?: string;
    /**
     * Title Suffix. Must be populated if ADVANCED_TITLE is on.
     * @type {number}
     * @memberof PersonNameType
     */
    nameTitleSuffix?: number;
    /**
     * Envelope Greeting of the profile
     * @type {string}
     * @memberof PersonNameType
     */
    envelopeGreeting?: string;
    /**
     * Salutation of the profile
     * @type {string}
     * @memberof PersonNameType
     */
    salutation?: string;
    /**
     *
     * @type {PersonNameTypeType}
     * @memberof PersonNameType
     */
    nameType?: PersonNameTypeType;
    /**
     * Language identification.
     * @type {string}
     * @memberof PersonNameType
     */
    language?: string;
}
/**
 * Check if a given object implements the PersonNameType interface.
 */
export declare function instanceOfPersonNameType(value: object): boolean;
export declare function PersonNameTypeFromJSON(json: any): PersonNameType;
export declare function PersonNameTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PersonNameType;
export declare function PersonNameTypeToJSON(value?: PersonNameType | null): any;
